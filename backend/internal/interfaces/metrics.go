package interfaces

import (
	"pinger/internal/models"
)

type IMetricsRepository interface {
	Create(metric models.LatencyMetric) error
	GetByMonitorId(monitorId string) ([]models.LatencyMetric, error)
}
