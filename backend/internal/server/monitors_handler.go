package server

import (
	"errors"
	"pinger/internal/dto"
	"pinger/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (s *Server) create(c *gin.Context) {
	var req dto.CreateMonitorDto
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", formatCreateMonitorValidationError(err))
		return
	}

	monitor, err := s.MonitorService.Create(&req)
	if err != nil {
		utils.BadRequestResponse(c, "failed to create monitor", err)
		return
	}

	utils.CreatedResponse(c, "Monitor created successfully", monitor)
}

func (s *Server) findAll(c *gin.Context) {
	monitors, err := s.MonitorService.FindAll(c.Request.Context())
	if err != nil {
		utils.InternalServerErrorResponse(c, "failed to list monitors", err)
		return
	}

	utils.SuccessResponse(c, "Monitors listed successfully", monitors)
}

func formatCreateMonitorValidationError(err error) error {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return errors.New("Invalid JSON request body")
	}

	for _, fieldError := range validationErrors {
		switch fieldError.Field() {
		case "URL":
			if fieldError.Tag() == "required" {
				return errors.New("The url cannot be empty")
			}
			return errors.New("The url must be valid")
		case "IntervalSeconds":
			return errors.New("The Interval should be at least 30 secondos long")
		}
	}

	return errors.New("Invalid request data")
}
