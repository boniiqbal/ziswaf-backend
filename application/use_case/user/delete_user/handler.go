package delete_user

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteUserHandler struct {
	request    infrastructure.Request
	repository infrastructure.UserRepository
}

func NewDeleteUserHandler(request infrastructure.Request, usRepo infrastructure.UserRepository) DeleteUserHandler {
	return DeleteUserHandler{
		request:    request,
		repository: usRepo,
	}
}

func (handler *DeleteUserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteUserByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Akun berhasil dihapus", true))
}
