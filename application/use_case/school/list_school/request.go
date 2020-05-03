package list_school

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	SchoolFilterRequest struct {
		Search      string `json:"search"`
		Regency     string `json:"regency"`
		Province    string `json:"province"`
		Sort        string `json:"sort"`
		Transaction string `json:"transaction"`
		Page        string `json:"Page"`
		Limit       string `json:"limit"`
		Detail      string `json:"detail"`
	}
)

func RequestMapper(req SchoolFilterRequest) domain.SchoolFilter {
	return domain.SchoolFilter{
		Search:      req.Search,
		Regency:     req.Regency,
		Province:    req.Province,
		Sort:        req.Sort,
		Transaction: req.Transaction,
		Page:        req.Page,
		Limit:       req.Limit,
		Detail:      req.Detail,
	}
}
