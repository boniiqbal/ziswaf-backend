package create_school

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreateSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewCreateSchoolHandler(request infrastructure.Request, schoolRepo infrastructure.SchoolRepository) CreateSchoolHandler {
	return CreateSchoolHandler{
		request:    request,
		repository: schoolRepo,
	}
}

func (handler *CreateSchoolHandler) CreateSchoolHandler(c *gin.Context) {
	request := CreateSchoolRequest{}
	acc, _ := c.Get("UserId")

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

	school, err := handler.repository.CreateSchool(ctx, RequestMapper(request, acc.(uint64)))
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(school), "Sekolah berhasil dibuat", true))
}
