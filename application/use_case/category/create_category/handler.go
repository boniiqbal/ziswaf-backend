package create_category

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreateCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewCreateCategoryHandler(request infrastructure.Request, ctrgRepo infrastructure.CategoryRepository) CreateCategoryHandler {
	return CreateCategoryHandler{
		request:    request,
		repository: ctrgRepo,
	}
}

func (handler *CreateCategoryHandler) CreateCategoryHandler(c *gin.Context) {
	request := CreateCategoryRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)

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

	cat, err := handler.repository.CreateCategory(ctx, RequestMapper(request, accountID))
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(cat), "Berhasil membuat kategori baru", true))
}
