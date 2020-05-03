package list_province

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListProvincesHandler struct {
	request    infrastructure.Request
	repository infrastructure.ProvinceRepository
}

func NewListProvincesHandler(request infrastructure.Request, provRepo infrastructure.ProvinceRepository) ListProvincesHandler {
	return ListProvincesHandler{
		request:    request,
		repository: provRepo,
	}
}

func (handler *ListProvincesHandler) ListProvincesHandler(c *gin.Context) {
	search := c.Query("search")
	school := c.Query("filter[school]")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	province := handler.repository.GetListProvince(ctx, search, school)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(province), "List province Success", true))
	return
}
