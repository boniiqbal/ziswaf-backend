package export_kwitansi

import (
	_trx "ziswaf-backend/application/use_case/transaction/show_transaction"
	"ziswaf-backend/domain/entities"
	domain "ziswaf-backend/domain/entities"
)

func ResponseMapper(domain _trx.ShowTransactionResponseData) domain.Kwitansi {
	return entities.Kwitansi{
		ID:               domain.ID,
		Description:      domain.Description,
		DonorName:        domain.DonorName,
		DonorPhone:       domain.DonorPhone,
		DonorEmail:       domain.DonorEmail,
		DonorAddress:     domain.DonorAddress,
		DonorNPWP:        domain.DonorNPWP,
		DonorCategory:    domain.DonorCategory,
		DivisionName:     domain.DivisionName,
		Kwitansi:         domain.Kwitansi,
		Unit:             domain.Unit,
		City:             domain.City,
		Category:         domain.Category,
		Total:            domain.Total,
		ItemType:         domain.ItemType,
		ItemCategory:     domain.ItemCategory,
		RefNumber:        domain.RefNumber,
		Quantity:         domain.Quantity,
		Status:           domain.Status,
		CreatedAt:        domain.CreatedAt,
		CashDescription:  domain.CashDescription,
		GoodDescription:  domain.GoodDescription,
		GoodStatus:       domain.GoodStatus,
		StatmentCategory: domain.StatmentCategory,
		CreatedBy:        domain.CreatedBy,
		SchoolName:       domain.SchoolName,
	}
}
