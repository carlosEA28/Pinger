package workers

import (
	"context"
	"pinger/internal/interfaces"

	"github.com/google/uuid"
)

type SchedulerConfig struct {
	monitorRepository interfaces.IMonitorsRepository
	metricsRepository interfaces.IMetricsRepository
}

func Scheduler(ctx context.Context, intervalTime int, monitorId uuid.UUID) {

}
