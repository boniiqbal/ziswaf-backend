package update_user

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateUserRequest struct {
		Data struct {
			Name      string `json:"name"`
			Username  string `json:"username" validate:"min=6,max=20"`
			Password  string `json:"password" validate:"min=8"`
			UpdatedBy string `json:"updated_by"`
			Status    int    `json:"status"`
			Role      int    `json:"role"`
		}
	}
)

func ValidateRequest(req *UpdateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateUserRequest, password string, username string, accountID string) domain.User {
	return domain.User{
		Name:      req.Data.Name,
		Username:  username,
		Password:  password,
		UpdatedBy: accountID,
		Status:    req.Data.Status,
		Role:      req.Data.Role,
	}
}
