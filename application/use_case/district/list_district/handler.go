package list_district

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListDistrictHandler struct {
	request    infrastructure.Request
	repository infrastructure.DistrictRepository
}

func NewListDistrictHandler(request infrastructure.Request, regRepo infrastructure.DistrictRepository) ListDistrictHandler {
	return ListDistrictHandler{
		request:    request,
		repository: regRepo,
	}
}

func (handler *ListDistrictHandler) ListDistrictHandler(c *gin.Context) {
	search := c.Query("search")
	regencyID := c.Query("filter[regency_id]")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	district := handler.repository.GetListDistrict(ctx, search, regencyID)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(district), "List district Success", true))
	return
}
