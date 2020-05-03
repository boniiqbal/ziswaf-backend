package list_school

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

type ListSchoolHandler struct {
	request    infrastructure.Request
	repository infrastructure.SchoolRepository
}

func NewListSchoolHandler(request infrastructure.Request, regRepo infrastructure.SchoolRepository) ListSchoolHandler {
	return ListSchoolHandler{
		request:    request,
		repository: regRepo,
	}
}

func (handler *ListSchoolHandler) ListSchoolHandler(c *gin.Context) {
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

	filter := SchoolFilterRequest{
		Search:      c.Query("search"),
		Regency:     c.Query("filter[regency]"),
		Province:    c.Query("filter[province]"),
		Sort:        misc.SortQuery(c.Request.URL.Query()),
		Transaction: c.Query("filter[is_transaction]"),
		Page:        c.Query("paging[page]"),
		Limit:       c.DefaultQuery("paging[limit]", "10"),
		Detail:      c.Query("filter[detail]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	school, count := handler.repository.GetListSchool(ctx, RequestMapper(filter))

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(school), paginationResponse, "List school Success", true))
	return
}
