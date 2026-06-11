package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateMonitorDto struct {
	URL             string `json:"url" binding:"required,url"`
	IntervalSeconds int    `json:"intervalSeconds" binding:"required,min=30"`
	IsActive        bool   `json:"isActive"`
}

type MonitorResponseDto struct {
	ID              uuid.UUID `json:"id"`
	URL             string    `json:"url"`
	IntervalSeconds int       `json:"intervalSeconds"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
