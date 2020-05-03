package update_statement_category

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateStatementCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewUpdateStatementCategoryHandler(request infrastructure.Request, catRepo infrastructure.CategoryRepository) UpdateStatementCategoryHandler {
	return UpdateStatementCategoryHandler{
		request:    request,
		repository: catRepo,
	}
}

func (handler *UpdateStatementCategoryHandler) UpdateStatementCategoryHandler(c *gin.Context) {
	request := UpdateStatementCategoryRequest{}
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := misc.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	catData, err := handler.repository.UpdateStatementCategory(ctx, RequestMapper(request), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(catData), "Update keterangan donasi berhasil", true))
}
