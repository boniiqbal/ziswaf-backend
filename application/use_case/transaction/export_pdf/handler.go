package export_pdf

import (
	"context"
	"net/http"
	"time"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc/reports"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ExportPdfHandler struct {
	request    infrastructure.Request
	repository infrastructure.TransactionRepository
}

func NewExportPdfHandler(request infrastructure.Request, transactionRepo infrastructure.TransactionRepository) ExportPdfHandler {
	return ExportPdfHandler{
		request:    request,
		repository: transactionRepo,
	}
}

func (handler *ExportPdfHandler) ExportPdfHandler(c *gin.Context) {
	filter := PdfFilterRequest{
		SchoolID:  c.Query("filter[school_id]"),
		StartDate: c.Query("filter[start_date]"),
		EndDate:   c.Query("filter[end_date]"),
		Regency:   c.Query("filter[regency]"),
		Role:      c.Query("filter[role]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	export, err := handler.repository.ExportPdf(ctx, RequestMapper(filter))
	if err != nil {
		c.JSON(http.StatusNotFound, response.SetMessage(err.Error(), false))
		return
	}

	timeFormat := time.Now().Format("2006-01-02")

	pdf := reports.NewPdf(export)
	pdf.MonthlyReport()
	pdf.TypeDonationReport()
	pdf.DivisionReport()
	pdf.GoodsReport()
	pdf.CashReport()
	pdf.Footer()
	pdf.Output("Laporan Umum "+timeFormat, c)
}
