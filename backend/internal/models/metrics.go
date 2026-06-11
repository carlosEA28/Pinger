package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LatencyMetric struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	MonitorID      uuid.UUID `gorm:"type:uuid;not null;index"`
	Timestamp      time.Time `gorm:"not null;index"`
	ResponseTimeMs float64   `gorm:"not null"`
	StatusCode     int       `gorm:"not null"`
	DnsLookupMs    *float64
	TCPConnectMs   *float64
	TTFBMs         *float64
	IsUp           bool `gorm:"not null"`
}

func (m *LatencyMetric) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}

	return nil
}
