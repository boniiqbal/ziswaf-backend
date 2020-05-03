package export_pdf

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ExportPdfResponse struct {
		base.BaseResponse
		Data ExportPdfResponseData `json:"data"`
	}

	ExportPdfResponseData struct {
		DonationReport domain.ExportPdf `json:"donation_report"`
	}
)

func SetResponse(domain ExportPdfResponseData, message string, success bool) ExportPdfResponse {
	return ExportPdfResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.ExportPdf) ExportPdfResponseData {

	return ExportPdfResponseData{
		DonationReport: domain,
	}
}
