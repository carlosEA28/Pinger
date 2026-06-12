package models

import (
	"time"

	"github.com/google/uuid"
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
