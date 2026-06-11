package repository

import (
	"context"
	"pinger/internal/dto"
	"pinger/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormMonitorsRepository struct {
	db *gorm.DB
}

func NewGormMonitorsRepository(db *gorm.DB) *GormMonitorsRepository {
	return &GormMonitorsRepository{
		db: db,
	}
}

func (r *GormMonitorsRepository) Create(monitor *models.Monitor) error {
	return r.db.Create(monitor).Error
}

func (r *GormMonitorsRepository) FindAll(ctx context.Context) ([]models.Monitor, error) {
	var monitors []models.Monitor

	if err := r.db.WithContext(ctx).Find(&monitors).Error; err != nil {
		return nil, err
	}

	return monitors, nil
}

func (r *GormMonitorsRepository) FindById(id uuid.UUID) (models.Monitor, error) {
	var monitor models.Monitor

	if err := r.db.First(&monitor, "id = ?", id).Error; err != nil {
		return models.Monitor{}, err
	}

	return monitor, nil
}

func (r *GormMonitorsRepository) Update(id uuid.UUID, fields dto.CreateMonitorDto) error {
	return r.db.Model(&models.Monitor{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"url":              fields.URL,
			"interval_seconds": fields.IntervalSeconds,
			"is_active":        fields.IsActive,
		}).Error
}

func (r *GormMonitorsRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Monitor{}, "id = ?", id).Error
}

func (r *GormMonitorsRepository) FindAllActive(ctx context.Context) ([]models.Monitor, error) {
	var monitors []models.Monitor

	if err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&monitors).Error; err != nil {
		return nil, err
	}

	return monitors, nil
}
