package list_transaction

import (
	"time"
	"ziswaf-backend/application/misc"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListTransactionResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData   `json:"pagination"`
		Data       []ListTransactionResponseData `json:"data"`
	}

	ListTransactionResponseData struct {
		ID                uint64    `json:"id"`
		DonorName         string    `json:"donor_name"`
		City              string    `json:"city"`
		DivisionName      string    `json:"division_name"`
		CategoryName      string    `json:"category_name"`
		Description       string    `json:"description"`
		ItemType          string    `json:"item_type"`
		RefNum            string    `json:"ref_number"`
		ItemCategory      string    `json:"item_category"`
		Status            string    `json:"status"`
		Total             uint64    `json:"total"`
		Kwitansi          string    `json:"kwitansi"`
		Unit              string    `json:"unit"`
		Phone             string    `json:"phone"`
		CashDescription   string    `json:"cash_description"`
		GoodDescription   string    `json:"good_description"`
		StatementCategory string    `json:"statement_category"`
		DonorCategory     bool      `json:"donor_category"`
		QuantityGood      int32     `json:"quantity_good"`
		GoodStatus        int8      `json:"good_status"`
		CreatedBy         string    `json:"created_by"`
		UpdatedBy         string    `json:"updated_by"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
	}
)

func (res *ListTransactionResponse) AddDomain(transaction []domain.Transaction) {
	response := ListTransactionResponseData{}
	var itemType string
	var ref_number = "-"
	var item_category string
	var item_quantity int32
	var good_status int8
	var cash_desc string
	var good_desc string
	var sts string

	for _, trans := range transaction {

		switch trans.ItemType {
		case misc.CASH:
			itemType = "uang"
			item_quantity = 0
			good_status = 0
			good_desc = ""

			if len(trans.Cashes) > 0 {
				item_category = trans.Cashes[0].CashCategory.Name
				ref_number = trans.Cashes[0].RefNumber
				cash_desc = trans.Cashes[0].Description
			}
		case misc.GOODS:
			itemType = "barang"
			ref_number = "-"
			item_category = ""
			cash_desc = ""

			if len(trans.GoodsQuery) > 0 {
				item_category = trans.GoodsQuery[0].GoodsCategory.Name
				item_quantity = trans.GoodsQuery[0].Quantity
				good_status = trans.GoodsQuery[0].Status
				good_desc = trans.GoodsQuery[0].Description
			}
		}

		if trans.Status == misc.RECEIVED {
			sts = "Diterima"
		}

		response.ID = trans.ID
		response.City = trans.School.Regency.Name
		response.DonorName = trans.Donor.Name
		response.DivisionName = trans.Division.Name
		response.CategoryName = trans.Category.Name
		response.Description = trans.Description
		response.ItemType = itemType
		response.ItemCategory = item_category
		response.RefNum = ref_number
		response.Status = sts
		response.Total = trans.Total
		response.Kwitansi = trans.Kwitansi
		response.Unit = trans.School.Name
		response.Phone = trans.Donor.Phone
		response.DonorCategory = trans.Donor.IsCompany
		response.GoodStatus = good_status
		response.QuantityGood = item_quantity
		response.CashDescription = cash_desc
		response.GoodDescription = good_desc
		response.StatementCategory = trans.StatementCategory.Name
		response.CreatedBy = trans.CreatedBy
		response.UpdatedBy = trans.UpdatedBy
		response.CreatedAt = trans.CreatedAt
		response.UpdatedAt = trans.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListTransactionResponseData, pagination base.PaginationResponseData, message string, success bool) ListTransactionResponse {
	return ListTransactionResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.Transaction) []ListTransactionResponseData {
	response := ListTransactionResponse{}

	response.AddDomain(domain)
	return response.Data
}
