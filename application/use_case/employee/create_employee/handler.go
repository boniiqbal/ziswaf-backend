package create_employee

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/refactory-id/go-core-package/response"
)

type CreateEmployeeHandler struct {
	request    infrastructure.Request
	repository infrastructure.EmployeeRepository
}

func NewCreateEmployeeHandler(request infrastructure.Request, employeeRepo infrastructure.EmployeeRepository) CreateEmployeeHandler {
	return CreateEmployeeHandler{
		request:    request,
		repository: employeeRepo,
	}
}

func (handler *CreateEmployeeHandler) CreateEmployeeHandler(c *gin.Context) {
	employeeID, _ := c.Get("EmployeeId")
	role, _ := c.Get("Role")
	request := CreateEmployeeRequest{}
	maxSize := int64(9048000)

	var schoolID uint64

	// Validate Max file request
	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(fmt.Sprintf("Gambar terlalu besar, maksimal ukuran : %v", maxSize), false))
		return
	}

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

	// Validate Role User
	if role == 1 {
		schoolID = request.SchoolID
	} else {
		empID := strconv.FormatUint(employeeID.(uint64), 10)

		// Get employee
		employee, errEmp := handler.repository.GetEmployeeByID(ctx, empID)
		if errEmp != nil {
			c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(errEmp.Error(), false))
			return
		}
		schoolID = employee.School.ID
	}

	file, errFile := misc.UploadFile(c, "image")
	if errFile != nil {
		if !strings.Contains(errFile.Error(), "Empty File") {
			c.JSON(misc.GetErrorStatusCode(422), response.SetMessage(errFile.Error(), false))
			return
		}
	}

	employee, err := handler.repository.CreateEmployee(ctx, RequestMapper(request, schoolID), file.FileURL)
	if err != nil {
		c.JSON(misc.GetErrorStatusCode(400), response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, SetResponse(ResponseMapper(employee), "Petugas berhasil dibuat", true))
}
