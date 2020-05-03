package list_report_operator

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ListReportOperatorHandler struct {
	request    infrastructure.Request
	repository infrastructure.TransactionRepository
}

func NewListReportOperatorHandler(request infrastructure.Request, transactionRepo infrastructure.TransactionRepository) ListReportOperatorHandler {
	return ListReportOperatorHandler{
		request:    request,
		repository: transactionRepo,
	}
}

func (handler *ListReportOperatorHandler) ListReportOperatorHandler(c *gin.Context) {
	filter := ReportOperatorFilterRequest{
		SchoolID:  c.Query("filter[school_id]"),
		StartDate: c.Query("filter[start_date]"),
		EndDate:   c.Query("filter[end_date]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	report, err := handler.repository.GetListReportOperator(ctx, RequestMapper(filter))
	if err != nil {
		c.JSON(http.StatusNotFound, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(report), "List Report Success", true))
	return
}
