package list_student

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

type ListStudentHandler struct {
	request    infrastructure.Request
	repository infrastructure.StudentRepository
}

func NewListStudentHandler(request infrastructure.Request, studentRepo infrastructure.StudentRepository) ListStudentHandler {
	return ListStudentHandler{
		request:    request,
		repository: studentRepo,
	}
}

func (handler *ListStudentHandler) ListStudentHandler(c *gin.Context) {
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

	filter := StudentFilterRequest{
		Search:          c.Query("search"),
		Sort:            misc.SortQuery(c.Request.URL.Query()),
		Page:            c.Query("paging[page]"),
		Limit:           c.DefaultQuery("paging[limit]", "10"),
		SchoolID:        c.Query("filter[school_id]"),
		SosialStatus:    c.Query("filter[sosial_status]"),
		EducationStatus: c.Query("filter[education_status]"),
		AgeStart:        c.Query("filter[age_start]"),
		AgeEnd:          c.Query("filter[age_end]"),
		RegisteredStart: c.Query("filter[registered_start]"),
		RegisteredEnd:   c.Query("filter[registered_end]"),
		Province:        c.Query("filter[province]"),
		Regency:         c.Query("filter[regency]"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	student, count := handler.repository.GetListStudent(ctx, RequestMapper(filter))

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(student), paginationResponse, "List student Success", true))
	return
}
