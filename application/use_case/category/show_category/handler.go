package show_category

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowCategoryHandler struct {
	request    infrastructure.Request
	repository infrastructure.CategoryRepository
}

func NewShowCategoryHandler(request infrastructure.Request, prdRepo infrastructure.CategoryRepository) ShowCategoryHandler {
	return ShowCategoryHandler{
		request:    request,
		repository: prdRepo,
	}
}

func (handler *ShowCategoryHandler) ShowCategoryHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := handler.repository.GetCategoryByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(user), "Show user success", true))
}
