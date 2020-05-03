package create_transaction

import (
	"time"
	"ziswaf-backend/application/misc"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateTransactionRequest struct {
		Data struct {
			DonorID             uint64    `json:"donor_id"`
			DivisionID          uint16    `json:"division_id"`
			CategoryID          uint16    `json:"category_id"`
			StatementCategoryID uint64    `json:"statement_category_id"`
			Description         string    `json:"description"`
			DonationItem        int8      `json:"donation_item"`
			Kwitansi            string    `json:"kwitansi"`
			EmployeeID          uint64    `json:"employee_id"`
			CreatedAt           time.Time `json:"created_at"`
			Cash                Cash      `json:"cash"`
			Goods               Goods     `json:"goods"`
		}
	}

	Cash struct {
		TypeID     int8   `json:"type_id"`
		CategoryID int8   `json:"category_id"`
		Value      uint64 `json:"value"`
		RefNumber  string `json:"ref_number"`
		Status     int8   `json:"status"`
	}

	Goods struct {
		CategoryID  int8   `json:"category_id"`
		Description string `json:"description"`
		Quantity    int32  `json:"quantity"`
		Value       uint64 `json:"value"`
		Status      int8   `json:"status"`
	}
)

func RequestCashMapper(req CreateTransactionRequest, createdBy string, schoolID uint64, createAt time.Time) domain.Cash {
	var crtAt time.Time

	if createAt.IsZero() {
		crtAt = time.Now()
	} else {
		crtAt = createAt
	}

	return domain.Cash{
		TypeID:     req.Data.Cash.TypeID,
		CategoryID: req.Data.Cash.CategoryID,
		RefNumber:  req.Data.Cash.RefNumber,
		Status:     req.Data.Cash.Status,
		Transaction: domain.Transaction{
			DonorID:             req.Data.DonorID,
			DivisionID:          req.Data.DivisionID,
			CategoryID:          req.Data.CategoryID,
			StatementCategoryID: req.Data.StatementCategoryID,
			Description:         req.Data.Description,
			Kwitansi:            req.Data.Kwitansi,
			Total:               req.Data.Cash.Value,
			SchoolID:            schoolID,
			CreatedBy:           createdBy,
			UpdatedBy:           createdBy,
			CreatedAt:           crtAt,
			Status:              misc.RECEIVED,
		},
	}
}

func ValidateRequest(req *CreateTransactionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestGoodsMapper(req CreateTransactionRequest, createdBy string, schoolID uint64, createAt time.Time) domain.Goods {
	var crtAt time.Time

	if createAt.IsZero() {
		crtAt = time.Now()
	} else {
		crtAt = createAt
	}

	return domain.Goods{
		CategoryID:  req.Data.Goods.CategoryID,
		Quantity:    req.Data.Goods.Quantity,
		Status:      req.Data.Goods.Status,
		Description: req.Data.Goods.Description,
		Transaction: domain.Transaction{
			DonorID:             req.Data.DonorID,
			DivisionID:          req.Data.DivisionID,
			CategoryID:          req.Data.CategoryID,
			StatementCategoryID: req.Data.StatementCategoryID,
			Description:         req.Data.Description,
			Total:               req.Data.Goods.Value,
			SchoolID:            schoolID,
			CreatedBy:           createdBy,
			UpdatedBy:           createdBy,
			CreatedAt:           crtAt,
			Status:              misc.RECEIVED,
		},
	}
}
