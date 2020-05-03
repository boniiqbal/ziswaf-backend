package export_pdf

import (
	domain "ziswaf-backend/domain/entities"
)

type (
	PdfFilterRequest struct {
		SchoolID  string
		StartDate string
		EndDate   string
		Regency   string
		Role      string
	}
)

func RequestMapper(req PdfFilterRequest) domain.TransactionFilter {
	return domain.TransactionFilter{
		SchoolID:  req.SchoolID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Regency:   req.Regency,
		Role:      req.Role,
	}
}
