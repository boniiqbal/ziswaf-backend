package update_division

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateDivisionHandler struct {
	request    infrastructure.Request
	repository infrastructure.DivisionRepository
}

func NewUpdateDivisionHandler(request infrastructure.Request, divRepo infrastructure.DivisionRepository) UpdateDivisionHandler {
	return UpdateDivisionHandler{
		request:    request,
		repository: divRepo,
	}
}

func (handler *UpdateDivisionHandler) UpdateDivisionHandler(c *gin.Context) {
	request := UpdateDivisionRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)
	id := c.Param("id")

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

	divData, err := handler.repository.UpdateDivision(ctx, RequestMapper(request, accountID), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(divData), "Update divisi berhasil", true))
}
