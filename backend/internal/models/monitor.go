package models

import (
	"time"

	"github.com/google/uuid"
)

type Monitor struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	URL             string    `gorm:"not null"`
	IntervalSeconds int       `gorm:"not null;default:60"`
	IsActive        bool      `gorm:"not null;default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
