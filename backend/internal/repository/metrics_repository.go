package repository

import (
	"context"
	"pinger/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormMetricsRepository struct {
	db *gorm.DB
}

func NewGormMetricsRepository(db *gorm.DB) *GormMetricsRepository {
	return &GormMetricsRepository{db: db}
}

func (r *GormMetricsRepository) Create(ctx context.Context, metric models.LatencyMetric) error {
	return r.db.WithContext(ctx).Create(&metric).Error
}

func (r *GormMetricsRepository) GetByMonitorId(ctx context.Context, monitorId uuid.UUID) ([]models.LatencyMetric, error) {
	var metrics []models.LatencyMetric

	if err := r.db.WithContext(ctx).
		Where("monitor_id = ?", monitorId).
		Order("timestamp DESC").
		Find(&metrics).Error; err != nil {
		return nil, err
	}

	return metrics, nil
}
