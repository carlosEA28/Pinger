package interfaces

import (
	"context"
	"pinger/internal/models"

	"github.com/google/uuid"
)

type IMetricsRepository interface {
	Create(ctx context.Context, metric models.LatencyMetric) error
	GetByMonitorId(ctx context.Context, monitorId uuid.UUID) ([]models.LatencyMetric, error)
}
