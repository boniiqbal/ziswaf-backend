package update_school

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewUpdateSchoolHandler(request infrastructure.Request, schoolRepo infrastructure.SchoolRepository) UpdateSchoolHandler {
	return UpdateSchoolHandler{
		request:    request,
		repository: schoolRepo,
	}
}

func (handler *UpdateSchoolHandler) UpdateSchoolHandler(c *gin.Context) {
	request := UpdateSchoolRequest{}
	id := c.Param("id")
	accID, _ := c.Get("UserId")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err := handler.repository.GetSchoolByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
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

	schData, err := handler.repository.UpdateSchool(ctx, RequestMapper(request, accID.(uint64)), id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(schData), "Update sekolah berhasil", true))
}
