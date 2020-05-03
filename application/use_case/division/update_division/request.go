package update_division

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateDivisionRequest struct {
		Data struct {
			Name        string `json:"name" validation:"required"`
			Description string `json:"description"`
			Status      int    `json:"status"`
			UpdatedBy   string `json:"updated_by"`
		}
	}
)

func ValidateRequest(req *UpdateDivisionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateDivisionRequest, accountID string) domain.Division {
	return domain.Division{
		Name:        req.Data.Name,
		Description: req.Data.Description,
		Status:      req.Data.Status,
		UpdatedBy:   accountID,
	}
}
