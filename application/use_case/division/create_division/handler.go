package create_division

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreateDivisionHandler struct {
	request    infrastructure.Request
	repository infrastructure.DivisionRepository
}

func NewCreateDivisionHandler(request infrastructure.Request, prdRepo infrastructure.DivisionRepository) CreateDivisionHandler {
	return CreateDivisionHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *CreateDivisionHandler) CreateDivisionHandler(c *gin.Context) {
	request := CreateDivisionRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)

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

	dvs, err := handler.repository.CreateDivision(ctx, RequestMapper(request, accountID))
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(dvs), "Berhasil membuat divisi baru ", true))
}
