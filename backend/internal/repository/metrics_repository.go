package repository

import (
	"pinger/internal/models"

	"gorm.io/gorm"
)

type GormMetricsRepository struct {
	db *gorm.DB
}

func NewGormMetricsRepository(db *gorm.DB) *GormMetricsRepository {
	return &GormMetricsRepository{db: db}
}

func (r *GormMetricsRepository) Create(metric models.LatencyMetric) error {
	return r.db.Create(&metric).Error
}

func (r *GormMetricsRepository) GetByMonitorId(monitorId string) ([]models.LatencyMetric, error) {
	var metrics []models.LatencyMetric

	r.db.Where("monitor_id = ?", monitorId).Find(&metrics)

	return metrics, nil
}
