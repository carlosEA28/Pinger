package services

import (
	"context"
	"errors"
	"pinger/internal/dto"
	"pinger/internal/interfaces"
	"pinger/internal/models"
	"pinger/internal/workers"

	"github.com/google/uuid"
)

type MontiorsService struct {
	monitorsRepository interfaces.IMonitorsRepository
	metricsRepository  interfaces.IMetricsRepository
	pinger             *workers.Pinger
}

func NewMonitorsService(
	monitorsRepository interfaces.IMonitorsRepository,
	metricsRepository interfaces.IMetricsRepository,
) *MontiorsService {
	return &MontiorsService{
		monitorsRepository: monitorsRepository,
		metricsRepository:  metricsRepository,
		pinger:             workers.NewPinger(monitorsRepository, metricsRepository, nil),
	}
}

func (s *MontiorsService) Create(req *dto.CreateMonitorDto) (*dto.CreateMonitorDto, error) {

	monitor := models.Monitor{
		URL:             req.URL,
		IntervalSeconds: req.IntervalSeconds,
		IsActive:        req.IsActive,
	}

	if req.URL == "" {
		return nil, errors.New("The url cannot be empty")
	}

	if req.IntervalSeconds < 30 {
		return nil, errors.New("The Interval should be at least 30 secondos long")

	}

	if err := s.monitorsRepository.Create(&monitor); err != nil {
		return nil, err
	}

	return &dto.CreateMonitorDto{
		URL:             monitor.URL,
		IntervalSeconds: monitor.IntervalSeconds,
		IsActive:        monitor.IsActive,
	}, nil
}

func (s *MontiorsService) FindAll(ctx context.Context) ([]dto.MonitorResponseDto, error) {
	monitors, err := s.monitorsRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]dto.MonitorResponseDto, 0, len(monitors))
	for _, monitor := range monitors {
		response = append(response, monitorResponseDto(monitor))
	}

	return response, nil
}

func (s *MontiorsService) Update(id string, req *dto.UpdateMonitorDto) (*dto.MonitorResponseDto, error) {
	monitorID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid monitor id")
	}

	if req.URL == nil && req.RequestedIntervalSeconds() == nil && req.RequestedIsActive() == nil {
		return nil, errors.New("At least one field must be provided")
	}

	if req.URL != nil && *req.URL == "" {
		return nil, errors.New("The url cannot be empty")
	}

	if err := s.monitorsRepository.Update(monitorID, *req); err != nil {
		return nil, err
	}

	monitor, err := s.monitorsRepository.FindById(monitorID)
	if err != nil {
		return nil, err
	}

	response := monitorResponseDto(monitor)
	return &response, nil
}

func (s *MontiorsService) Delete(id string) error {
	monitorID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("Invalid monitor id")
	}

	return s.monitorsRepository.Delete(monitorID)
}

func (s *MontiorsService) Ping(ctx context.Context, id string) (*dto.MonitorResponseDto, error) {
	monitorID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid monitor id")
	}

	monitor, err := s.monitorsRepository.FindById(monitorID)
	if err != nil {
		return nil, err
	}

	s.pinger.Ping(ctx, monitor)

	monitor, err = s.monitorsRepository.FindById(monitorID)
	if err != nil {
		return nil, err
	}

	response := monitorResponseDto(monitor)
	return &response, nil
}

func (s *MontiorsService) Metrics(ctx context.Context, id string) ([]dto.LatencyMetricResponseDto, error) {
	monitorID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid monitor id")
	}

	if _, err := s.monitorsRepository.FindById(monitorID); err != nil {
		return nil, err
	}

	metrics, err := s.metricsRepository.GetByMonitorId(ctx, monitorID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.LatencyMetricResponseDto, 0, len(metrics))
	for _, metric := range metrics {
		response = append(response, latencyMetricResponseDto(metric))
	}

	return response, nil
}

func monitorResponseDto(monitor models.Monitor) dto.MonitorResponseDto {
	return dto.MonitorResponseDto{
		ID:              monitor.ID,
		URL:             monitor.URL,
		IntervalSeconds: monitor.IntervalSeconds,
		IsActive:        monitor.IsActive,
		LastCheckedAt:   monitor.LastCheckedAt,
		CreatedAt:       monitor.CreatedAt,
		UpdatedAt:       monitor.UpdatedAt,
	}
}

func latencyMetricResponseDto(metric models.LatencyMetric) dto.LatencyMetricResponseDto {
	return dto.LatencyMetricResponseDto{
		ID:             metric.ID,
		MonitorID:      metric.MonitorID,
		Timestamp:      metric.Timestamp,
		ResponseTimeMs: metric.ResponseTimeMs,
		StatusCode:     metric.StatusCode,
		DnsLookupMs:    metric.DnsLookupMs,
		TCPConnectMs:   metric.TCPConnectMs,
		TTFBMs:         metric.TTFBMs,
		IsUp:           metric.IsUp,
	}
}
