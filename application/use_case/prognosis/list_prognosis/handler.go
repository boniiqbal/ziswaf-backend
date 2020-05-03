package list_prognosis

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

type ListPrognosisHandler struct {
	request            infrastructure.Request
	repository         infrastructure.PrognosisRepository
	divisionRepository infrastructure.DivisionRepository
}

func NewListPrognosisHandler(request infrastructure.Request, prognosisRepo infrastructure.PrognosisRepository, divRepo infrastructure.DivisionRepository) ListPrognosisHandler {
	return ListPrognosisHandler{
		request:            request,
		repository:         prognosisRepo,
		divisionRepository: divRepo,
	}
}

func (handler *ListPrognosisHandler) ListPrognosisHandler(c *gin.Context) {
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

	filter := PrognosisFilterRequest{
		Search:     c.Query("search"),
		StartDate:  c.Query("filter[start_date]"),
		EndDate:    c.Query("filter[end_date]"),
		StartTotal: c.Query("filter[start_total]"),
		EndTotal:   c.Query("filter[end_total]"),
		DivisionID: c.Query("filter[division_id]"),
		Year:       c.Query("filter[year]"),
		Sort:       misc.SortQuery(c.Request.URL.Query()),
		Page:       c.Query("paging[page]"),
		Limit:      c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	prognosis, count := handler.repository.GetListPrognosis(ctx, RequestMapper(filter))
	division := handler.divisionRepository.GetListDivision(ctx)

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(prognosis, division), paginationResponse, "List prognosis Success", true))
	return
}
