package list_report

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	ReportFilterRequest struct {
		SchoolID  string
		StartDate string
		EndDate   string
		Regency   string
	}
)

func RequestMapper(req ReportFilterRequest) domain.TransactionFilter {
	return domain.TransactionFilter{
		SchoolID:  req.SchoolID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Regency:   req.Regency,
	}
}
