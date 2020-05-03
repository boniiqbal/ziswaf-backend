package list_donor

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	DonorFilterRequest struct {
		Search        string `json:"search"`
		Regency       string `json:"regency"`
		Status        string `json:"status"`
		SchoolID      string `json:"school_id"`
		DonorCategory string `json:"donor_category"`
		Sort          string `json:"sort"`
		Page          string `json:"Page"`
		Limit         string `json:"limit"`
	}
)

func RequestMapper(req DonorFilterRequest) domain.DonorFilter {
	return domain.DonorFilter{
		Search:        req.Search,
		Regency:       req.Regency,
		Status:        req.Status,
		SchoolID:      req.SchoolID,
		DonorCategory: req.DonorCategory,
		Sort:          req.Sort,
		Page:          req.Page,
		Limit:         req.Limit,
	}
}
