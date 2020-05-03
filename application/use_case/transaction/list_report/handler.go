package list_report

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ListReportHandler struct {
	request    infrastructure.Request
	repository infrastructure.TransactionRepository
}

func NewListReportHandler(request infrastructure.Request, transactionRepo infrastructure.TransactionRepository) ListReportHandler {
	return ListReportHandler{
		request:    request,
		repository: transactionRepo,
	}
}

func (handler *ListReportHandler) ListReportHandler(c *gin.Context) {
	filter := ReportFilterRequest{
		SchoolID:  c.Query("filter[school_id]"),
		StartDate: c.Query("filter[start_date]"),
		EndDate:   c.Query("filter[end_date]"),
		Regency:   c.Query("filter[regency]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	report, err := handler.repository.GetListReport(ctx, RequestMapper(filter))
	if err != nil {
		c.JSON(http.StatusNotFound, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(report), "List Report Success", true))
	return
}
