package delete_student

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteStudentHandler struct {
	request    infrastructure.Request
	repository infrastructure.StudentRepository
}

func NewDeleteStudentHandler(request infrastructure.Request, stuRepo infrastructure.StudentRepository) DeleteStudentHandler {
	return DeleteStudentHandler{
		request:    request,
		repository: stuRepo,
	}
}

func (handler *DeleteStudentHandler) DeleteStudentHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteStudentByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Siswa berhasil dihapus", true))
}
