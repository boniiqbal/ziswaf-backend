package list_donor

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

type ListDonorsHandler struct {
	request           infrastructure.Request
	repository        infrastructure.DonorRepository
	regencyRepository infrastructure.RegencyRepository
}

func NewListDonorsHandler(request infrastructure.Request, donorRepo infrastructure.DonorRepository, regRepo infrastructure.RegencyRepository) ListDonorsHandler {
	return ListDonorsHandler{
		request:           request,
		repository:        donorRepo,
		regencyRepository: regRepo,
	}
}

func (handler *ListDonorsHandler) ListDonorsHandler(c *gin.Context) {
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

	filter := DonorFilterRequest{
		Search:        c.Query("search"),
		Regency:       c.Query("filter[regency]"),
		Status:        c.Query("filter[status]"),
		SchoolID:      c.Query("filter[school_id]"),
		DonorCategory: c.Query("filter[donor_category]"),
		Sort:          misc.SortQuery(c.Request.URL.Query()),
		Page:          c.Query("paging[page]"),
		Limit:         c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	donor, count := handler.repository.GetListDonor(ctx, RequestMapper(filter))
	regency := handler.regencyRepository.GetListRegency(ctx, "", "", "")

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(donor, regency), paginationResponse, "List donor Success", true))
	return
}
