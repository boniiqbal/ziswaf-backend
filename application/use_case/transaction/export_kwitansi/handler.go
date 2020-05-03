package export_kwitansi

import (
	"context"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
	"ziswaf-backend/application/misc/reports"
	_trx "ziswaf-backend/application/use_case/transaction/show_transaction"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ExportKwitansiHandler struct {
	request       infrastructure.Request
	repository    infrastructure.TransactionRepository
	usrRepository infrastructure.UserRepository
}

func NewExportKwitansiHandler(request infrastructure.Request, trxRepo infrastructure.TransactionRepository, usrRepo infrastructure.UserRepository) ExportKwitansiHandler {
	return ExportKwitansiHandler{
		request:       request,
		repository:    trxRepo,
		usrRepository: usrRepo,
	}
}

func (handler *ExportKwitansiHandler) ExportKwitansiHandler(c *gin.Context) {
	id := c.Param("id")
	var officerName string

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	trx, err := handler.repository.ShowTransaction(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(404), response.SetMessage(err.Error(), false))
		return
	}

	// Get User
	user, err := handler.usrRepository.GetUserById(ctx, trx.CreatedBy)
	if err != nil {
		officerName = "-"
	} else {
		officerName = user.Employee.Name
	}

	data := ResponseMapper(_trx.ResponseMapper(trx, officerName))

	pdf := reports.NewKwitansiPdf(data)
	pdf.Header()
	pdf.Body()
	pdf.Output("Detail Transaksi - "+data.DonorName, c)

	// c.JSON(http.StatusOK, SetResponse(ResponseMapper(trx, officerName), "Show transaction success", true))
}
