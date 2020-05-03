package show_student

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowStudentHandler struct {
	request    infrastructure.Request
	repository infrastructure.StudentRepository
}

func NewShowStudentHandler(request infrastructure.Request, stuRepo infrastructure.StudentRepository) ShowStudentHandler {
	return ShowStudentHandler{
		request:    request,
		repository: stuRepo,
	}
}

func (handler *ShowStudentHandler) ShowStudentHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	student, err := handler.repository.GetStudentByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(student), "Show student success", true))
}
