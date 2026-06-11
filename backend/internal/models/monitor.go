package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Monitor struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	URL             string    `gorm:"not null"`
	IntervalSeconds int       `gorm:"not null;default:60"`
	IsActive        bool      `gorm:"not null;default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (m *Monitor) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}

	return nil
}
