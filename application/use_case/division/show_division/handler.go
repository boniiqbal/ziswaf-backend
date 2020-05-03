package show_division

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowDivisionHandler struct {
	request    infrastructure.Request
	repository infrastructure.DivisionRepository
}

func NewShowDivisionHandler(request infrastructure.Request, divisionRepo infrastructure.DivisionRepository) ShowDivisionHandler {
	return ShowDivisionHandler{
		request:    request,
		repository: divisionRepo,
	}
}

func (handler *ShowDivisionHandler) ShowDivisionHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	division, err := handler.repository.GetDivisionByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(division), "Show division success", true))
}
