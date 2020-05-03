package list_report_donation

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListReportDonationResponse struct {
		base.BaseResponse
		Data ListReportDonationResponseData `json:"data"`
	}

	ListReportDonationResponseData struct {
		Total          uint64                   `json:"total"`
		TotalRowCount  int                      `json:"total_row_count"`
		DonationReport []donationReportResponse `json:"donation_report"`
	}

	donationReportResponse struct {
		StatementCategoryID uint64 `json:"statement_category_id"`
		Name                string `json:"name"`
		Total               uint64 `json:"total"`
		TotalPercent        uint64 `json:"total_percent"`
	}
)

func SetResponse(domain ListReportDonationResponseData, message string, success bool) ListReportDonationResponse {
	return ListReportDonationResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.ReportStatementCategory) ListReportDonationResponseData {
	donReport := donationReportResponse{}
	reportDonation := ListReportDonationResponseData{}

	reportDonation.Total = domain.TotalCash
	reportDonation.TotalRowCount = domain.TotalRowCount

	for _, v := range domain.Data {
		donReport.StatementCategoryID = v.StatementCategoryID
		donReport.Name = v.Name
		donReport.Total = v.Total
		donReport.TotalPercent = v.TotalPercent

		reportDonation.DonationReport = append(reportDonation.DonationReport, donReport)
	}

	return reportDonation
}
