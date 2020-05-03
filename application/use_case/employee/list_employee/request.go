package list_employee

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	EmployeeFilterRequest struct {
		Search        string `json:"search"`
		RegisterStart string `json:"start_date"`
		RegisterEnd   string `json:"end_date"`
		SchoolID      string `json:"school_id"`
		Status        string `json:"status"`
		Province      string `json:"province"`
		Regency       string `json:"regency"`
		Sort          string `json:"sort"`
		Page          string `json:"page"`
		Limit         string `json:"limit"`
	}
)

func RequestMapper(req EmployeeFilterRequest) domain.EmployeeFilter {
	return domain.EmployeeFilter{
		Search:        req.Search,
		RegisterStart: req.RegisterStart,
		RegisterEnd:   req.RegisterEnd,
		SchoolID:      req.SchoolID,
		Status:        req.Status,
		Province:      req.Province,
		Regency:       req.Regency,
		Sort:          req.Sort,
		Page:          req.Page,
		Limit:         req.Limit,
	}
}
