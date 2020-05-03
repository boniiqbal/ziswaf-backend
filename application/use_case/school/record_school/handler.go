package record_school

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
)

type RecordSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewRecordSchoolHandler(request infrastructure.Request, schoolRepo infrastructure.SchoolRepository) RecordSchoolHandler {
	return RecordSchoolHandler{
		request:    request,
		repository: schoolRepo,
	}
}

func (handler *RecordSchoolHandler) RecordSchoolHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	records, err := handler.repository.GetRecordSchool(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(records), "Show record of school success", true))
}
