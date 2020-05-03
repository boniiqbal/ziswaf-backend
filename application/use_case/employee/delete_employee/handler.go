package delete_employee

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type DeleteEmployeeHandler struct {
	request    infrastructure.Request
	repository infrastructure.EmployeeRepository
}

func NewDeleteEmployeeHandler(request infrastructure.Request, empRepo infrastructure.EmployeeRepository) DeleteEmployeeHandler {
	return DeleteEmployeeHandler{
		request:    request,
		repository: empRepo,
	}
}

func (handler *DeleteEmployeeHandler) DeleteEmployeeHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := handler.repository.DeleteEmployeeByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse("Petugas berhasil dihapus", true))
}
