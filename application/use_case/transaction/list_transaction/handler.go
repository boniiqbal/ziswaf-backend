package list_transaction

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
	_donor "ziswaf-backend/application/use_case/donor/list_donor"

	"github.com/gin-gonic/gin"
	base "github.com/refactory-id/go-core-package/response"
)

type ListTransactionHandler struct {
	request         infrastructure.Request
	repository      infrastructure.TransactionRepository
	repositoryDonor infrastructure.DonorRepository
}

func NewListTransactionHandler(request infrastructure.Request, transactionRepo infrastructure.TransactionRepository, donorRepo infrastructure.DonorRepository) ListTransactionHandler {
	return ListTransactionHandler{
		request:         request,
		repository:      transactionRepo,
		repositoryDonor: donorRepo,
	}
}

func (handler *ListTransactionHandler) ListTransactionHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("paging[page]"))
	limit, _ := strconv.Atoi(c.DefaultQuery("paging[limit]", "10"))
	regency := c.Query("filter[regency]")
	donaturStatus := c.Query("filter[status_donatur]")
	search := c.Query("search")
	sortType := "sort[" + strings.TrimSpace(misc.SortQuery(c.Request.URL.Query())) + "]"

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

	filterDonor := _donor.DonorFilterRequest{
		Regency: regency,
		Status:  donaturStatus,
		Search:  search,
		Page:    "",
		Limit:   "10",
	}

	filter := TransactionFilterRequest{
		Search:            search,
		StartDate:         c.Query("filter[start_date]"),
		EndDate:           c.Query("filter[end_date]"),
		StartTotal:        c.Query("filter[start_total]"),
		EndTotal:          c.Query("filter[end_total]"),
		Regency:           regency,
		DonaturStatus:     donaturStatus,
		DonaturCategory:   c.Query("filter[donor_category]"),
		StatementCategory: c.Query("filter[statement_category]"),
		DonationType:      c.Query("filter[category_donation]"),
		CategoryType:      c.Query("filter[category_type]"),
		SchoolID:          c.Query("filter[school_id]"),
		DonationCategory:  c.Query("filter[category_id]"),
		DonationSource:    c.Query("filter[division_id]"),
		Sort:              misc.SortQuery(c.Request.URL.Query()),
		SortType:          c.Query(sortType),
		Page:              c.Query("paging[page]"),
		Limit:             c.DefaultQuery("paging[limit]", "10"),
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	donor, _ := handler.repositoryDonor.GetListDonor(ctx, _donor.RequestMapper(filterDonor))
	transaction, count := handler.repository.GetListTransaction(ctx, RequestMapper(filter), donor)

	paginationResponse := base.PaginationResponseData{
		CurrentPage:    page,
		PageCount:      int(math.Ceil(float64(count) / float64(limit))),
		PageSize:       limit,
		RowCount:       count,
		FirstRowOnPage: firstRowOnPage,
		LastRowOnPage:  lastRowOnPage,
	}

	c.JSON(http.StatusOK, SetResponse(ResponseMapper(transaction), paginationResponse, "List transaction Success", true))
	return
}
