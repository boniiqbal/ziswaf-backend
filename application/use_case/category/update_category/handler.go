package update_category

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewUpdateCategoryHandler(request infrastructure.Request, catRepo infrastructure.CategoryRepository) UpdateCategoryHandler {
	return UpdateCategoryHandler{
		request:    request,
		repository: catRepo,
	}
}

func (handler *UpdateCategoryHandler) UpdateCategoryHandler(c *gin.Context) {
	request := UpdateCategoryRequest{}
	acc, _ := c.Get("UserId")
	accountID := strconv.FormatUint(acc.(uint64), 10)
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

	catData, err := handler.repository.UpdateCategory(ctx, RequestMapper(request, accountID), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(catData), "Update kategori berhasil", true))
}
