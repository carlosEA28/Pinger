package interfaces

import (
	"context"
	"pinger/internal/dto"
	"pinger/internal/models"

	"github.com/google/uuid"
)

type IMonitorsRepository interface {
	Create(monitors *models.Monitor) error
	FindAll(ctx context.Context) ([]models.Monitor, error)
	FindById(id uuid.UUID) (models.Monitor, error)
	Update(id uuid.UUID, fields dto.UpdateMonitorDto) error
	Delete(id uuid.UUID) error
	FindAllActive(ctx context.Context) ([]models.Monitor, error)
}
