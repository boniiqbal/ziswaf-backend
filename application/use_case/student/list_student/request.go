package list_student

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	StudentFilterRequest struct {
		Search          string `json:"search"`
		Sort            string `json:"sort"`
		SchoolID        string `json:"school_id"`
		SosialStatus    string `json:"sosial_status"`
		EducationStatus string `json:"education_status"`
		AgeStart        string `json:"age_start"`
		AgeEnd          string `json:"age_end"`
		RegisteredStart string `json:"registered_start"`
		RegisteredEnd   string `json:"registered_end"`
		Province        string `json:"province"`
		Regency         string `json:"regency"`
		Page            string `json:"page"`
		Limit           string `json:"limit"`
	}
)

func RequestMapper(req StudentFilterRequest) domain.StudentFilter {
	return domain.StudentFilter{
		Search:          req.Search,
		Sort:            req.Sort,
		SchoolID:        req.SchoolID,
		SosialStatus:    req.SosialStatus,
		EducationStatus: req.EducationStatus,
		AgeStart:        req.AgeStart,
		AgeEnd:          req.AgeEnd,
		RegisteredStart: req.RegisteredStart,
		RegisteredEnd:   req.RegisteredEnd,
		Province:        req.Province,
		Regency:         req.Regency,
		Page:            req.Page,
		Limit:           req.Limit,
	}
}
