package services

import (
	"errors"
	"pinger/internal/dto"
	"pinger/internal/interfaces"
	"pinger/internal/models"
)

type MontiorsService struct {
	monitorsRepository interfaces.IMonitorsRepository
}

func NewMonitorsService(monitorsRepository interfaces.IMonitorsRepository) *MontiorsService {
	return &MontiorsService{
		monitorsRepository: monitorsRepository,
	}
}

func (s *MontiorsService) Create(req *dto.CreateMonitorDto) (*dto.CreateMonitorDto, error) {

	monitor := models.Monitor{
		URL:             req.URL,
		IntervalSeconds: req.IntervalSeconds,
		IsActive:        req.IsActive,
	}

	if req.URL == "" {
		return nil, errors.New("The url cannot be empty")
	}

	if req.IntervalSeconds < 30 {
		return nil, errors.New("The Interval should be at least 30 secondos long")

	}

	if err := s.monitorsRepository.Create(&monitor); err != nil {
		return nil, err
	}

	return &dto.CreateMonitorDto{
		URL:             monitor.URL,
		IntervalSeconds: monitor.IntervalSeconds,
		IsActive:        monitor.IsActive,
	}, nil
}
