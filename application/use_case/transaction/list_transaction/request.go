package list_transaction

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	TransactionFilterRequest struct {
		Search            string `json:"search"`
		StartDate         string `json:"start_date"`
		EndDate           string `json:"end_date"`
		StartTotal        string `json:"start_total"`
		EndTotal          string `json:"end_total"`
		Regency           string `json:"regency"`
		DonaturStatus     string `json:"status_donatur"`
		DonaturCategory   string `json:"donatur_category"`
		StatementCategory string `json:"statement_category"`
		CategoryID        string `json:"category_id"`
		CategoryType      string `json:"category_type"`
		SchoolID          string `json:"school_id"`
		DonationCategory  string `json:"donation_category"`
		DonationSource    string `json:"donation_source"`
		DonationType      string `json:"donation_type"`
		Sort              string `json:"sort"`
		SortType          string `json:"sort_type"`
		Page              string `json:"page"`
		Limit             string `json:"limit"`
	}
)

func RequestMapper(req TransactionFilterRequest) domain.TransactionFilter {
	return domain.TransactionFilter{
		Search:            req.Search,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		StartTotal:        req.StartTotal,
		EndTotal:          req.EndTotal,
		Regency:           req.Regency,
		DonaturStatus:     req.DonaturStatus,
		Sort:              req.Sort,
		SortType:          req.SortType,
		Page:              req.Page,
		Limit:             req.Limit,
		DonaturCategory:   req.DonaturCategory,
		StatementCategory: req.StatementCategory,
		SchoolID:          req.SchoolID,
		CategoryID:        req.DonationCategory,
		DivisionID:        req.DonationSource,
		DonationType:      req.DonationType,
		CategoryType:      req.CategoryType,
	}
}
