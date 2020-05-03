package list_report_donation

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	ReportDonationFilterRequest struct {
		SchoolID          string
		StartDate         string
		EndDate           string
		Regency           string
		DivisionID        string
		DonaturCategory   string
		StatementCategory string
		CategoryID        string
		CategoryType      string
	}
)

func RequestMapper(req ReportDonationFilterRequest) domain.TransactionFilter {
	return domain.TransactionFilter{
		SchoolID:          req.SchoolID,
		StartDate:         req.StartDate,
		EndDate:           req.EndDate,
		Regency:           req.Regency,
		DivisionID:        req.DivisionID,
		DonaturCategory:   req.DonaturCategory,
		StatementCategory: req.StatementCategory,
		CategoryID:        req.CategoryID,
		CategoryType:      req.CategoryType,
	}
}
