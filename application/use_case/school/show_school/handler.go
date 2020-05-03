package show_school

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewShowSchoolHandler(request infrastructure.Request, schoolRepo infrastructure.SchoolRepository) ShowSchoolHandler {
	return ShowSchoolHandler{
		request:    request,
		repository: schoolRepo,
	}
}

func (handler *ShowSchoolHandler) ShowSchoolHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	school, err := handler.repository.GetSchoolByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(school), "Show school success", true))
}
