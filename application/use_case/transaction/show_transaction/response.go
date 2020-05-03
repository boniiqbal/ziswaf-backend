package show_transaction

import (
	"time"
	"ziswaf-backend/application/misc"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowTransactionResponse struct {
		base.BaseResponse
		Data ShowTransactionResponseData `json:"data"`
	}

	ShowTransactionResponseData struct {
		ID               uint64    `json:"id"`
		Description      string    `json:"description"`
		DonorName        string    `json:"donor_name"`
		DonorAddress     string    `json:"donor_address,omitempty"`
		DonorPhone       string    `json:"donor_phone"`
		DonorEmail       string    `json:"donor_email"`
		DonorNPWP        int64     `json:"donor_npwp"`
		DivisionName     string    `json:"division_name"`
		Unit             string    `json:"unit"`
		City             string    `json:"city"`
		Kwitansi         string    `json:"kwitansi"`
		Category         string    `json:"category"`
		StatmentCategory string    `json:"statement_category,omitempty"`
		SchoolName       string    `json:"school_name"`
		Total            uint64    `json:"total"`
		ItemType         string    `json:"item_type"`
		ItemCategory     string    `json:"item_category"`
		RefNumber        string    `json:"ref_number"`
		Quantity         int32     `json:"quantity"`
		Status           int8      `json:"status"`
		CreatedAt        time.Time `json:"created_at"`
		CashDescription  string    `json:"cash_description"`
		GoodDescription  string    `json:"good_description"`
		DonorCategory    bool      `json:"donor_category"`
		GoodStatus       int8      `json:"good_status"`
		CreatedBy        string    `json:"created_by"`
	}
)

func SetResponse(res ShowTransactionResponseData, message string, success bool) ShowTransactionResponse {
	return ShowTransactionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: res,
	}
}

func ResponseMapper(domain domain.Transaction, officerName string) ShowTransactionResponseData {
	var itemType string
	var refNumber = "-"
	var itemCategory string
	var qty int32
	var goodStatus int8
	var cashDesc string
	var goodDesc string

	switch domain.ItemType {
	case misc.CASH:
		itemType = "Uang"
		goodDesc = ""
		itemCategory = domain.Cashes[0].CashCategory.Name
		refNumber = domain.Cashes[0].RefNumber
		cashDesc = domain.Cashes[0].Description
		goodStatus = 0
		qty = 0
	case misc.GOODS:
		itemType = "Barang"
		cashDesc = ""
		refNumber = "-"
		itemCategory = domain.GoodsQuery[0].GoodsCategory.Name
		qty = domain.GoodsQuery[0].Quantity
		goodDesc = domain.GoodsQuery[0].Description
		goodStatus = domain.GoodsQuery[0].Status
	}

	return ShowTransactionResponseData{
		ID:               domain.ID,
		Description:      domain.Description,
		DonorName:        domain.Donor.Name,
		DonorPhone:       domain.Donor.Phone,
		DonorEmail:       domain.Donor.Email,
		DonorAddress:     domain.Donor.Address,
		DonorNPWP:        domain.Donor.Npwp,
		DonorCategory:    domain.Donor.IsCompany,
		DivisionName:     domain.Division.Name,
		Kwitansi:         domain.Kwitansi,
		Unit:             domain.School.Name,
		City:             domain.School.Regency.Name,
		SchoolName:       domain.School.Name,
		Category:         domain.Category.Name,
		Total:            domain.Total,
		ItemType:         itemType,
		ItemCategory:     itemCategory,
		RefNumber:        refNumber,
		Quantity:         qty,
		Status:           domain.Status,
		CreatedAt:        domain.CreatedAt,
		CashDescription:  cashDesc,
		GoodDescription:  goodDesc,
		GoodStatus:       goodStatus,
		StatmentCategory: domain.StatementCategory.Name,
		CreatedBy:        officerName,
	}
}
