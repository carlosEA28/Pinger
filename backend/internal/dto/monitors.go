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

type UpdateMonitorDto struct {
	URL                  *string `json:"url" binding:"omitempty,url"`
	IntervalSeconds      *int    `json:"intervalSeconds" binding:"omitempty,min=30"`
	IntervalSecondsSnake *int    `json:"interval_seconds" binding:"omitempty,min=30"`
	IsActive             *bool   `json:"isActive"`
	IsActiveSnake        *bool   `json:"is_active"`
}

func (d UpdateMonitorDto) RequestedIntervalSeconds() *int {
	if d.IntervalSeconds != nil {
		return d.IntervalSeconds
	}

	return d.IntervalSecondsSnake
}

func (d UpdateMonitorDto) RequestedIsActive() *bool {
	if d.IsActive != nil {
		return d.IsActive
	}

	return d.IsActiveSnake
}

type MonitorResponseDto struct {
	ID              uuid.UUID `json:"id"`
	URL             string    `json:"url"`
	IntervalSeconds int       `json:"intervalSeconds"`
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
