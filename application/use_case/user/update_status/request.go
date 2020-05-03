package update_status

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateStatusRequest struct {
		Data struct {
			UpdatedBy string `json:"updated_by"`
			Status    int    `json:"status"`
		}
	}
)

func ValidateRequest(req *UpdateStatusRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateStatusRequest, accountID string) domain.User {
	return domain.User{
		UpdatedBy: accountID,
		Status:    req.Data.Status,
	}
}
