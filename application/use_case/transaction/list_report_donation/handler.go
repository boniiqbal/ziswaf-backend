package list_report_donation

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ListReportDonationHandler struct {
	request    infrastructure.Request
	repository infrastructure.TransactionRepository
	repoCategory infrastructure.CategoryRepository
}

func NewListReportDonationHandler(request infrastructure.Request, transactionRepo infrastructure.TransactionRepository, repoCat infrastructure.CategoryRepository) ListReportDonationHandler {
	return ListReportDonationHandler{
		request:    request,
		repository: transactionRepo,
		repoCategory: repoCat,
	}
}

func (handler *ListReportDonationHandler) ListReportDonationHandler(c *gin.Context) {
	filter := ReportDonationFilterRequest{
		SchoolID:          c.Query("filter[school_id]"),
		StartDate:         c.Query("filter[start_date]"),
		EndDate:           c.Query("filter[end_date]"),
		Regency:           c.Query("filter[regency]"),
		DivisionID:        c.Query("filter[division_id]"),
		DonaturCategory:   c.Query("filter[donatur_category]"),
		StatementCategory: c.Query("filter[statement_category]"),
		CategoryID:        c.Query("filter[category_id]"),
		CategoryType:      c.Query("filter[category_type]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	stCategories := handler.repoCategory.GetListStatementCategory(ctx, "1")

	report, err := handler.repository.GetListDonationReport(ctx, RequestMapper(filter), stCategories)
	if err != nil {
		c.JSON(http.StatusNotFound, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(report), "List Report Success", true))
	return
}
