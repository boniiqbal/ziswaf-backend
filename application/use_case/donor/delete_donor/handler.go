package delete_donor

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteDonorHandler struct {
	request    infrastructure.Request
	repository infrastructure.DonorRepository
}

func NewDeleteDonorHandler(request infrastructure.Request, donRepo infrastructure.DonorRepository) DeleteDonorHandler {
	return DeleteDonorHandler{
		request:    request,
		repository: donRepo,
	}
}

func (handler *DeleteDonorHandler) DeleteDonorHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteDonorByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Donatur berhasil dihapus", true))
}
