package workers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
	"pinger/internal/interfaces"
	"pinger/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	defaultPingTimeout      = 30 * time.Second
	defaultSchedulerBackoff = 30 * time.Second
	minSchedulerInterval    = time.Second
)

type SchedulerConfig struct {
	SchedulerInterval time.Duration
	HTTPTimeout       time.Duration
}

type Scheduler struct {
	monitorRepository interfaces.IMonitorsRepository
	metricsRepository interfaces.IMetricsRepository
	pinger            *Pinger
	backoff           time.Duration
}

func NewScheduler(
	monitorRepository interfaces.IMonitorsRepository,
	metricsRepository interfaces.IMetricsRepository,
) *Scheduler {
	return NewSchedulerWithConfig(monitorRepository, metricsRepository, SchedulerConfig{})
}

func NewSchedulerWithConfig(
	monitorRepository interfaces.IMonitorsRepository,
	metricsRepository interfaces.IMetricsRepository,
	cfg SchedulerConfig,
) *Scheduler {
	httpTimeout := cfg.HTTPTimeout
	if httpTimeout <= 0 {
		httpTimeout = defaultPingTimeout
	}

	backoff := cfg.SchedulerInterval
	if backoff <= 0 {
		backoff = defaultSchedulerBackoff
	}

	return &Scheduler{
		monitorRepository: monitorRepository,
		metricsRepository: metricsRepository,
		pinger: NewPinger(
			monitorRepository,
			metricsRepository,
			&http.Client{Timeout: httpTimeout},
		),
		backoff: backoff,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	for {
		nextRunIn := s.runCycle(ctx)
		timer := time.NewTimer(nextRunIn)

		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
		}
	}
}

func (s *Scheduler) runCycle(ctx context.Context) time.Duration {
	monitors, err := s.monitorRepository.FindAllActive(ctx)
	if err != nil {
		log.Printf("scheduler: failed to load active monitors: %v", err)
		return s.backoff
	}

	now := time.Now()
	nextRunIn := s.backoff
	var wg sync.WaitGroup

	for _, monitor := range monitors {
		interval := monitorInterval(monitor)
		if !isReadyToPing(monitor, now) {
			nextRunIn = minDuration(nextRunIn, nextPingIn(monitor, now))
			continue
		}

		nextRunIn = minDuration(nextRunIn, interval)
		wg.Add(1)
		go func(monitor models.Monitor) {
			defer func() {
				if recovered := recover(); recovered != nil {
					log.Printf("scheduler: recovered from pinger panic for monitor %s: %v", monitor.ID, recovered)
				}
				wg.Done()
			}()

			s.pinger.Ping(ctx, monitor)
		}(monitor)
	}

	wg.Wait()
	return nextRunIn
}

func isReadyToPing(monitor models.Monitor, now time.Time) bool {
	if monitor.LastCheckedAt == nil {
		return true
	}

	return now.Sub(*monitor.LastCheckedAt) >= monitorInterval(monitor)
}

func nextPingIn(monitor models.Monitor, now time.Time) time.Duration {
	if monitor.LastCheckedAt == nil {
		return 0
	}

	nextRunIn := monitorInterval(monitor) - now.Sub(*monitor.LastCheckedAt)
	if nextRunIn < minSchedulerInterval {
		return minSchedulerInterval
	}

	return nextRunIn
}

func monitorInterval(monitor models.Monitor) time.Duration {
	if monitor.IntervalSeconds <= 0 {
		return minSchedulerInterval
	}

	return time.Duration(monitor.IntervalSeconds) * time.Second
}

func minDuration(a, b time.Duration) time.Duration {
	if b < a {
		return b
	}

	return a
}

type Pinger struct {
	monitorRepository interfaces.IMonitorsRepository
	metricsRepository interfaces.IMetricsRepository
	client            *http.Client
}

func NewPinger(
	monitorRepository interfaces.IMonitorsRepository,
	metricsRepository interfaces.IMetricsRepository,
	client *http.Client,
) *Pinger {
	if client == nil {
		client = &http.Client{Timeout: defaultPingTimeout}
	}

	return &Pinger{
		monitorRepository: monitorRepository,
		metricsRepository: metricsRepository,
		client:            client,
	}
}

func (p *Pinger) Ping(ctx context.Context, monitor models.Monitor) {
	startedAt := time.Now()
	trace := &pingTrace{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, monitor.URL, nil)
	if err != nil {
		p.saveResult(ctx, monitor.ID, startedAt, 0, 0, nil, nil, nil, false)
		return
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace.clientTrace()))

	resp, err := p.client.Do(req)
	if err != nil {
		p.saveResult(ctx, monitor.ID, startedAt, 0, elapsedMs(startedAt, time.Now()), trace.dnsLookupMs(), trace.tcpConnectMs(), trace.ttfbMs(startedAt), false)
		return
	}
	defer resp.Body.Close()

	finishedAt := time.Now()
	responseTimeMs := elapsedMs(startedAt, finishedAt)
	isUp := resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest

	p.saveResult(ctx, monitor.ID, startedAt, resp.StatusCode, responseTimeMs, trace.dnsLookupMs(), trace.tcpConnectMs(), trace.ttfbMs(startedAt), isUp)
}

func (p *Pinger) saveResult(
	ctx context.Context,
	monitorID uuid.UUID,
	timestamp time.Time,
	statusCode int,
	responseTimeMs float64,
	dnsLookupMs *float64,
	tcpConnectMs *float64,
	ttfbMs *float64,
	isUp bool,
) {
	if err := p.metricsRepository.Create(ctx, models.LatencyMetric{
		ID:             uuid.New(),
		MonitorID:      monitorID,
		Timestamp:      timestamp,
		ResponseTimeMs: responseTimeMs,
		StatusCode:     statusCode,
		DnsLookupMs:    dnsLookupMs,
		TCPConnectMs:   tcpConnectMs,
		TTFBMs:         ttfbMs,
		IsUp:           isUp,
	}); err != nil {
		log.Printf("pinger: failed to save latency metric for monitor %s: %v", monitorID, err)
	}

	if err := p.monitorRepository.UpdateLastCheckedAt(ctx, monitorID, timestamp); err != nil {
		log.Printf("pinger: failed to update last_checked_at for monitor %s: %v", monitorID, err)
	}
}

type pingTrace struct {
	dnsStart      time.Time
	dnsDone       time.Time
	connectStart  time.Time
	connectDone   time.Time
	firstResponse time.Time
}

func (t *pingTrace) clientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			t.dnsStart = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			t.dnsDone = time.Now()
		},
		ConnectStart: func(_, _ string) {
			t.connectStart = time.Now()
		},
		ConnectDone: func(_, _ string, _ error) {
			t.connectDone = time.Now()
		},
		GotFirstResponseByte: func() {
			t.firstResponse = time.Now()
		},
	}
}

func (t *pingTrace) dnsLookupMs() *float64 {
	if t.dnsStart.IsZero() || t.dnsDone.IsZero() {
		return nil
	}

	value := elapsedMs(t.dnsStart, t.dnsDone)
	return &value
}

func (t *pingTrace) tcpConnectMs() *float64 {
	if t.connectStart.IsZero() || t.connectDone.IsZero() {
		return nil
	}

	value := elapsedMs(t.connectStart, t.connectDone)
	return &value
}

func (t *pingTrace) ttfbMs(startedAt time.Time) *float64 {
	if t.firstResponse.IsZero() {
		return nil
	}

	value := elapsedMs(startedAt, t.firstResponse)
	return &value
}

func elapsedMs(start, end time.Time) float64 {
	return end.Sub(start).Seconds() * 1000
}
