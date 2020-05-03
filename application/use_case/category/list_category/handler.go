package list_category

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"

	"github.com/gin-gonic/gin"
)

type ListCategoriesHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewListCategoriesHandler(request infrastructure.Request, prdRepo infrastructure.CategoryRepository) ListCategoriesHandler {
	return ListCategoriesHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *ListCategoriesHandler) ListCategoriesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	category := handler.repository.GetListCategory(ctx)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(category), "List category Success", true))
	return
}
