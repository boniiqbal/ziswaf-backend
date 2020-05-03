package create_division

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateDivisionRequest struct {
		Data struct {
			Name        string `json:"name" validation:"required"`
			Description string `json:"description"`
			Status      int    `json:"status"`
			CreatedBy   string `json:"created_by"`
			UpdatedBy   string `json:"updated_by"`
		}
	}
)

func ValidateRequest(req *CreateDivisionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateDivisionRequest, accountID string) domain.Division {
	return domain.Division{
		Name:        req.Data.Name,
		Description: req.Data.Description,
		Status:      req.Data.Status,
		CreatedBy:   accountID,
		UpdatedBy:   accountID,
	}
}
