package create_statement_category

import (
	validator "gopkg.in/go-playground/validator.v9"
	domain "ziswaf-backend/domain/entities"
)

type (
	CreateStatementCategoryRequest struct {
		Data struct {
			StatementCategoryName string `json:"jenis_kategori" validation:"required"`
			CategoryID            uint64 `json:"kategory_id"`
		}
	}
)

func ValidateRequest(req *CreateStatementCategoryRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateStatementCategoryRequest) domain.StatementCategory {
	return domain.StatementCategory{
		CategoryID: req.Data.CategoryID,
		Name:       req.Data.StatementCategoryName,
	}
}
