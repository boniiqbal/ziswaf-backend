package list_report_operator

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	ReportOperatorFilterRequest struct {
		SchoolID  string
		StartDate string
		EndDate   string
	}
)

func RequestMapper(req ReportOperatorFilterRequest) domain.TransactionFilter {
	return domain.TransactionFilter{
		SchoolID:  req.SchoolID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
}
