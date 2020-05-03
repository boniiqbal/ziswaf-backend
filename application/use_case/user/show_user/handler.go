package show_user

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowUserHandler struct {
	request       infrastructure.Request
	prdRepository infrastructure.UserRepository
}

func NewShowUserHandler(request infrastructure.Request, prdRepo infrastructure.UserRepository) ShowUserHandler {
	return ShowUserHandler{
		request:       request,
		prdRepository: prdRepo,
	}
}

func (handler *ShowUserHandler) ShowUserHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := handler.prdRepository.GetUserById(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(user), "Show user success", true))
}
