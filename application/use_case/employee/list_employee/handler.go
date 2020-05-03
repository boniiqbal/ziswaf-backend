package list_employee

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"

	"github.com/gin-gonic/gin"
	base "github.com/refactory-id/go-core-package/response"
)

type ListEmployeesHandler struct {
	request    infrastructure.Request
	repository infrastructure.EmployeeRepository
}

func NewListEmployeesHandler(request infrastructure.Request, employeeRepo infrastructure.EmployeeRepository) ListEmployeesHandler {
	return ListEmployeesHandler{
		request:    request,
		repository: employeeRepo,
	}
}

func (handler *ListEmployeesHandler) ListEmployeesHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("paging[page]"))
	limit, _ := strconv.Atoi(c.DefaultQuery("paging[limit]", "10"))

	var (
		firstRowOnPage int
		lastRowOnPage  int
	)

	lastRowOnPage = page * limit
	if page == 1 {
		firstRowOnPage = page
	} else {
		firstRowOnPage = lastRowOnPage - limit + 1
	}

	filter := EmployeeFilterRequest{
		Search:  c.Query("search"),
		Regency: c.Query("filter[regency]"),
		Sort:    misc.SortQuery(c.Request.URL.Query()),
		Status:  c.Query("filter[status]"),
		RegisterStart: c.Query("filter[register_start]"),
		RegisterEnd: c.Query("filter[register_end]"),
		SchoolID: c.Query("filter[school_id]"),
		Province: c.Query("filter[province]"),
		Page:    c.Query("paging[page]"),
		Limit:   c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	employee, count := handler.repository.GetListEmployee(ctx, RequestMapper(filter))

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(employee), paginationResponse, "List employee Success", true))
	return
}
