package repository

import (
	"context"
	"pinger/internal/dto"
	"pinger/internal/models"
	"time"

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

func (r *GormMonitorsRepository) Update(id uuid.UUID, fields dto.UpdateMonitorDto) error {
	updates := map[string]interface{}{}

	if fields.URL != nil {
		updates["url"] = *fields.URL
	}

	if intervalSeconds := fields.RequestedIntervalSeconds(); intervalSeconds != nil {
		updates["interval_seconds"] = *intervalSeconds
	}

	if isActive := fields.RequestedIsActive(); isActive != nil {
		updates["is_active"] = *isActive
	}

	return r.db.Model(&models.Monitor{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormMonitorsRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Monitor{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *GormMonitorsRepository) FindAllActive(ctx context.Context) ([]models.Monitor, error) {
	var monitors []models.Monitor

	if err := r.db.WithContext(ctx).Where("is_active = ?", true).Find(&monitors).Error; err != nil {
		return nil, err
	}

	return monitors, nil
}

func (r *GormMonitorsRepository) UpdateLastCheckedAt(ctx context.Context, id uuid.UUID, t time.Time) error {
	return r.db.WithContext(ctx).
		Model(&models.Monitor{}).
		Where("id = ?", id).
		Update("last_checked_at", t).Error
}
