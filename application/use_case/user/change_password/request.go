package change_password

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	ChangeUserRequest struct {
		Data struct {
			Password        string `json:"password" validate:"required"`
			NewPassword     string `json:"new_password" validate:"required,min=8"`
			ConfirmPassword string `json:"confirm_password" validate:"required"`
		}
	}
)

func ValidateRequest(req *ChangeUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(password string, accountID string) domain.User {
	return domain.User{
		Password:  password,
		UpdatedBy: accountID,
	}
}
