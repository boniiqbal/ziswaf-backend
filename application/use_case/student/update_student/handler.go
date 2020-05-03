package update_student

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type UpdateStudentHandler struct {
	request    infrastructure.Request
	repository infrastructure.StudentRepository
}

func NewUpdateStudentHandler(request infrastructure.Request, stuRepo infrastructure.StudentRepository) UpdateStudentHandler {
	return UpdateStudentHandler{
		request:    request,
		repository: stuRepo,
	}
}

func (handler *UpdateStudentHandler) UpdateStudentHandler(c *gin.Context) {
	request := UpdateStudentRequest{}
	id := c.Param("id")
	maxSize := int64(9048000)

	// Validate Max file request
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(fmt.Sprintf("Gambar terlalu besar, maksimal ukuran : %v", maxSize), false))
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err = handler.repository.GetStudentByID(ctx, id)
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

	file, errFile := misc.UploadFile(c, "image")
	if errFile != nil {
		if !strings.Contains(errFile.Error(), "Empty File") {
			c.JSON(misc.GetErrorStatusCode(422), response.SetMessage(errFile.Error(), false))
			return
		}
	}

	stuData, err := handler.repository.UpdateStudent(ctx, RequestMapper(request), id, file.FileURL)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}
	c.JSON(misc.GetErrorStatusCode(200), SetResponse(ResponseMapper(stuData), "Update siswa berhasil", true))
}
