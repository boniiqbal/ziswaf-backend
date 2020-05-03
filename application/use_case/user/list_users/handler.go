package list_users

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
	_employee "ziswaf-backend/application/use_case/employee/list_employee"

	base "github.com/refactory-id/go-core-package/response"

	"github.com/gin-gonic/gin"
)

type ListUsersHandler struct {
	request       infrastructure.Request
	prdRepository infrastructure.UserRepository
	empRepository infrastructure.EmployeeRepository
}

func NewListUsersHandler(request infrastructure.Request, prdRepo infrastructure.UserRepository, empRepo infrastructure.EmployeeRepository) ListUsersHandler {
	return ListUsersHandler{
		request:       request,
		prdRepository: prdRepo,
		empRepository: empRepo,
	}
}

func (handler *ListUsersHandler) ListUsersHandler(c *gin.Context) {
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

	filterEmployee := _employee.EmployeeFilterRequest{
		SchoolID: c.Query("filter[school_id]"),
		Search:   c.Query("search"),
		Page:     "",
		Limit:    "10",
	}

	filter := UserFilterRequest{
		Search:       c.Query("search"),
		Sort:         misc.SortQuery(c.Request.URL.Query()),
		SchoolID:     c.Query("filter[school_id]"),
		Role:         c.Query("filter[role]"),
		CreatedStart: c.Query("filter[created_start]"),
		CreatedEnd:   c.Query("filter[created_end]"),
		LoginStart:   c.Query("filter[login_start]"),
		LoginEnd:     c.Query("filter[login_end]"),
		Page:         c.Query("paging[page]"),
		Limit:        c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	employee, _ := handler.empRepository.GetListEmployee(ctx, _employee.RequestMapper(filterEmployee))
	users, count := handler.prdRepository.ListUsers(ctx, RequestMapper(filter), employee)

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(users), paginationResponse, "List Users Success", true))
	return
}
