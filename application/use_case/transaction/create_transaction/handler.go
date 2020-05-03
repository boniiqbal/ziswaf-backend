package create_transaction

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"

	domain "ziswaf-backend/domain/entities"
)

type CreateTransactionHandler struct {
	request       infrastructure.Request
	repository    infrastructure.TransactionRepository
	empRepository infrastructure.EmployeeRepository
	usrRepository infrastructure.UserRepository
}

func NewCreateTransactionHandler(request infrastructure.Request, trxRepo infrastructure.TransactionRepository, empRepo infrastructure.EmployeeRepository, usrRepo infrastructure.UserRepository) CreateTransactionHandler {
	return CreateTransactionHandler{
		request:       request,
		repository:    trxRepo,
		empRepository: empRepo,
		usrRepository: usrRepo,
	}
}

func (handler *CreateTransactionHandler) CreateTransactionHandler(c *gin.Context) {
	var item interface{}
	var trxID uint64
	var dvsID uint16
	var employeeID uint64
	var createdBy string
	var createdAt time.Time
	var errTrx error

	acc, _ := c.Get("UserId")
	emplID, _ := c.Get("EmployeeId")
	inputType := c.Query("type")
	accountID := strconv.FormatUint(acc.(uint64), 10)
	request := CreateTransactionRequest{}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := misc.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	if request.Data.DivisionID == 1 && inputType != "" {
		employeeID = request.Data.EmployeeID
		createdAt = request.Data.CreatedAt

		// Get User
		user, err := handler.usrRepository.GetUserByEmployeeId(ctx, employeeID)
		if err != nil {
			c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
			return
		}
		userID := strconv.FormatUint(user.ID, 10)

		createdBy = userID
	} else if request.Data.DivisionID == 1 && inputType == "" {
		employeeID = emplID.(uint64)
		createdBy = accountID
	} else {
		createdBy = accountID
		employeeID = emplID.(uint64)
	}

	empID := strconv.FormatUint(employeeID, 10)

	// Get Employee
	employee, err := handler.empRepository.GetEmployeeByID(ctx, empID)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	if request.Data.DonationItem != 1 && request.Data.DonationItem != 2 {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage("Item donasi hanya bisa 1 atau 2", false))
		return
	}

	if request.Data.DonationItem == 1 {
		item, errTrx = handler.repository.CreateCash(ctx, RequestCashMapper(request, createdBy, employee.SchoolID, createdAt))
		trx := item.(domain.Cash)
		trxID = trx.Transaction.ID
		dvsID = trx.Transaction.DivisionID
	}

	if request.Data.DonationItem == 2 {
		item, errTrx = handler.repository.CreateGoods(ctx, RequestGoodsMapper(request, createdBy, employee.SchoolID, createdAt))
		trx := item.(domain.Goods)
		trxID = trx.Transaction.ID
		dvsID = trx.Transaction.DivisionID
	}

	if trxID == 0 {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errTrx.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(trxID, dvsID), "Transaksi berhasil dibuat", true))
}
