package dto

type CreateMonitorDto struct {
	URL             string `json:"url" binding:"required,url"`
	IntervalSeconds int    `json:"intervalSeconds" binding:"required,min=30"`
	IsActive        bool   `json:"isActive"`
}
