package update_status

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateStatusHandler struct {
	request    infrastructure.Request
	repository infrastructure.UserRepository
}

func NewUpdateStatusHandler(request infrastructure.Request, prdRepo infrastructure.UserRepository) UpdateStatusHandler {
	return UpdateStatusHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *UpdateStatusHandler) UpdateStatusHandler(c *gin.Context) {
	request := UpdateStatusRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Get user by id
	_, err := handler.repository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
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

	_, err = handler.repository.UpdateUser(ctx, RequestMapper(request, accountID), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	usData, err := handler.repository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(usData), "Update user berhasil", true))
}
