package list_village

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListVillageHandler struct {
	request    infrastructure.Request
	repository infrastructure.VillageRepository
}

func NewListVillageHandler(request infrastructure.Request, vilRepo infrastructure.VillageRepository) ListVillageHandler {
	return ListVillageHandler{
		request:    request,
		repository: vilRepo,
	}
}

func (handler *ListVillageHandler) ListVillageHandler(c *gin.Context) {
	search := c.Query("search")
	districtID := c.Query("filter[district_id]")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	village := handler.repository.GetListVillage(ctx, search, districtID)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(village), "List district Success", true))
	return
}
