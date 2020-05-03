package login

import (
	"ziswaf-backend/application/misc"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	LoginRequest struct {
		Data struct {
			Username  string `json:"username" validate:"required"`
			Password  string `json:"password" validate:"required"`
			UserID    uint64 `json:"user_id"`
			Token     string `json:"token"`
			ExpiredAt int64  `json:"expire_at"`
		}
	}
)

func ValidateRequest(req *LoginRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(userId uint64, signed string) domain.AccessToken {
	return domain.AccessToken{
		UserID:    userId,
		Token:     signed,
		ExpiredAt: misc.GetExpiryTime(misc.EXPIRED_AT),
	}
}
