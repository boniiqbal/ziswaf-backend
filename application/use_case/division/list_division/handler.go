package list_division

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListDivisionsHandler struct {
	request    infrastructure.Request
	repository infrastructure.DivisionRepository
}

func NewListDivisionsHandler(request infrastructure.Request, divRepo infrastructure.DivisionRepository) ListDivisionsHandler {
	return ListDivisionsHandler{
		request:    request,
		repository: divRepo,
	}
}

func (handler *ListDivisionsHandler) ListDivisionsHandler(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	division := handler.repository.GetListDivision(ctx)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(division), "List division Success", true))
	return
}
