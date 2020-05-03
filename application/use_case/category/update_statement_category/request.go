package update_statement_category

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateStatementCategoryRequest struct {
		Data struct {
			Name       string `json:"name"`
			CategoryID uint64 `json:"category_id"`
		}
	}
)

func ValidateRequest(req *UpdateStatementCategoryRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateStatementCategoryRequest) domain.StatementCategory {
	return domain.StatementCategory{
		CategoryID: req.Data.CategoryID,
		Name:       req.Data.Name,
	}
}
