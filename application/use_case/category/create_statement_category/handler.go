package create_statement_category

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
)

type CreateStatementCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewCreateStatementCategoryHandler(request infrastructure.Request, ctgrRepo infrastructure.CategoryRepository) CreateStatementCategoryHandler {
	return CreateStatementCategoryHandler{
		request:    request,
		repository: ctgrRepo,
	}
}

func (handler *CreateStatementCategoryHandler) CreateStatementCategoryHandler(c *gin.Context) {
	request := CreateStatementCategoryRequest{}

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

	ctgr, err := handler.repository.CreateStatementCategory(ctx, RequestMapper(request))
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(http.StatusInternalServerError), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(ctgr), "Jenis Kategori baru berhasil dibuat", true))
}
