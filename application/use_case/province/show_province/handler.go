package show_province

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowProvinceHandler struct {
	request    infrastructure.Request
	repository infrastructure.ProvinceRepository
}

func NewShowProvinceHandler(request infrastructure.Request, provRepo infrastructure.ProvinceRepository) ShowProvinceHandler {
	return ShowProvinceHandler{
		request:    request,
		repository: provRepo,
	}
}

func (handler *ShowProvinceHandler) ShowProvinceHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	province, err := handler.repository.GetProvinceByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(province), "Show province success", true))
}
