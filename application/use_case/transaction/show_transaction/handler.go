package show_transaction

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type ShowTransactionHandler struct {
	request       infrastructure.Request
	repository    infrastructure.TransactionRepository
	usrRepository infrastructure.UserRepository
}

func NewShowTransactionHandler(request infrastructure.Request, trxRepo infrastructure.TransactionRepository, usrRepo infrastructure.UserRepository) ShowTransactionHandler {
	return ShowTransactionHandler{
		request:       request,
		repository:    trxRepo,
		usrRepository: usrRepo,
	}
}

func (handler *ShowTransactionHandler) ShowTransactionHandler(c *gin.Context) {
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

	if trx.DivisionID == 1 {
		// Get User
		user, err := handler.usrRepository.GetUserById(ctx, trx.CreatedBy)
		if err != nil {
			officerName = "-"
		} else {
			officerName = user.Employee.Name
		}

	} else {
		officerName = trx.CreatedBy
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(trx, officerName), "Show transaction success", true))
}
