package list_regency

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListRegenciesHandler struct {
	request    infrastructure.Request
	repository infrastructure.RegencyRepository
}

func NewListRegenciesHandler(request infrastructure.Request, regRepo infrastructure.RegencyRepository) ListRegenciesHandler {
	return ListRegenciesHandler{
		request:    request,
		repository: regRepo,
	}
}

func (handler *ListRegenciesHandler) ListRegenciesHandler(c *gin.Context) {
	search := c.Query("search")
	provinceID := c.Query("filter[province_id]")
	school := c.Query("filter[school]")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	regency := handler.repository.GetListRegency(ctx, search, provinceID, school)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(regency), "List regency Success", true))
	return
}
