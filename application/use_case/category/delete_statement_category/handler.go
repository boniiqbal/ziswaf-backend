package delete_statement_category

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteStatementCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewDeleteStatementCategoryHandler(request infrastructure.Request, cat infrastructure.CategoryRepository) DeleteStatementCategoryHandler {
	return DeleteStatementCategoryHandler{
		request:    request,
		repository: cat,
	}
}

func (handler *DeleteStatementCategoryHandler) DeleteStatementCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteStatementCategoryByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Keterangan donasi berhasil dihapus", true))
}
