package create_user

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateUserRequest struct {
		Data struct {
			Username   string `json:"username" validate:"required,min=6,max=20"`
			Password   string `json:"password" validate:"required,min=8"`
			EmployeeID uint64 `json:"employee_id" validate:"required"`
			Name       string `json:"name"`
			CreatedBy  string `json:"created_by"`
			UpdatedBy  string `json:"updated_by"`
			Status     int    `json:"status"`
			Role       int    `json:"role"`
		}
	}
)

func ValidateRequest(req *CreateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateUserRequest, password string, username string, accountID string) domain.User {
	return domain.User{
		Username:   username,
		Password:   password,
		EmployeeID: req.Data.EmployeeID,
		Name:       req.Data.Name,
		CreatedBy:  accountID,
		UpdatedBy:  accountID,
		Status:     req.Data.Status,
		Role:       req.Data.Role,
	}
}
