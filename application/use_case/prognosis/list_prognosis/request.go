package list_prognosis

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	PrognosisFilterRequest struct {
		Search     string `json:"search"`
		StartDate  string `json:"start_date"`
		EndDate    string `json:"end_date"`
		StartTotal string `json:"start_total"`
		EndTotal   string `json:"end_total"`
		Sort       string `json:"sort"`
		DivisionID string `json:"division_id"`
		Year       string `json:"year"`
		Page       string `json:"page"`
		Limit      string `json:"limit"`
	}
)

func RequestMapper(req PrognosisFilterRequest) domain.TransactionGoalFilter {
	return domain.TransactionGoalFilter{
		Search:     req.Search,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		StartTotal: req.StartTotal,
		EndTotal:   req.EndTotal,
		Sort:       req.Sort,
		DivisionID: req.DivisionID,
		Year:       req.Year,
		Page:       req.Page,
		Limit:      req.Limit,
	}
}
