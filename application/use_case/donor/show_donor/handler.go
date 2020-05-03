package show_donor

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowDonorHandler struct {
	request           infrastructure.Request
	repository        infrastructure.DonorRepository
	regencyRepository infrastructure.RegencyRepository
}

func NewShowDonorHandler(request infrastructure.Request, donorRepo infrastructure.DonorRepository, regRepo infrastructure.RegencyRepository) ShowDonorHandler {
	return ShowDonorHandler{
		request:           request,
		repository:        donorRepo,
		regencyRepository: regRepo,
	}
}

func (handler *ShowDonorHandler) ShowDonorHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	donor, err := handler.repository.GetDonorByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	regency := handler.regencyRepository.GetListRegency(ctx, "", "", "")

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(donor, regency), "Show donor success", true))
}
