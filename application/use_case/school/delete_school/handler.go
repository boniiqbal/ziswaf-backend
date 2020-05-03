package delete_school

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewDeleteSchoolHandler(request infrastructure.Request, schRepo infrastructure.SchoolRepository) DeleteSchoolHandler {
	return DeleteSchoolHandler{
		request:    request,
		repository: schRepo,
	}
}

func (handler *DeleteSchoolHandler) DeleteSchoolHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteSchoolByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Ma'Had berhasil dihapus", true))
}
