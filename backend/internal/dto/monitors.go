package dto

type CreateMonitorDto struct {
	URL             string
	IntervalSeconds int
	IsActive        bool
}
