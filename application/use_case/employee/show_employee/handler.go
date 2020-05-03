package show_employee

import (
	"context"
	"net/http"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"

	"github.com/refactory-id/go-core-package/response"
)

type ShowEmployeeHandler struct {
	request           infrastructure.Request
	repository        infrastructure.EmployeeRepository
}

func NewShowEmployeeHandler(request infrastructure.Request, employeeRepo infrastructure.EmployeeRepository) ShowEmployeeHandler {
	return ShowEmployeeHandler{
		request:           request,
		repository:        employeeRepo,
	}
}

func (handler *ShowEmployeeHandler) ShowEmployeeHandler(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	employee, err := handler.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(employee), "Show employee success", true))
}
