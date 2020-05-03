package list_statement_category

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ziswaf-backend/application/infrastructure"
)

type ListStatementCategoriesHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewListStatementCategoriesHandler(request infrastructure.Request, ctgrRepo infrastructure.CategoryRepository) ListStatementCategoriesHandler {
	return ListStatementCategoriesHandler{
		request:    request,
		repository: ctgrRepo,
	}
}

func (handler *ListStatementCategoriesHandler) ListStatementCategoriesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	trx := c.Query("filter[is_transaction]")

	stCategories := handler.repository.GetListStatementCategory(ctx, trx)

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(stCategories), "List Keterangan Kategori", true))
}
