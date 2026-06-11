package interfaces

import (
	"context"
	"pinger/internal/dto"
	"pinger/internal/models"

	"github.com/google/uuid"
)

type IMonitorsRepository interface {
	Create(ctx context.Context, monitors *dto.CreateMonitorDto) error
	FindAll(ctx context.Context) ([]models.Monitor, error)
	FindById(ctx context.Context, id uuid.UUID) (models.Monitor, error)
	Update(ctx context.Context, id uuid.UUID, fields dto.CreateMonitorDto) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindAllActive(ctx context.Context) ([]models.Monitor, error)
}
