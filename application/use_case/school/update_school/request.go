package update_school

import (
	"strconv"
	"time"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateSchoolRequest struct {
		Data struct {
			Name        string    `json:"name" validate:"required"`
			Phone       string    `json:"phone" validate:"numeric"`
			Email       string    `json:"email" validate:"email"`
			Address     string    `json:"address"`
			PosCode     int       `json:"pos_code"`
			Description string    `json:"description"`
			ProvinceID  uint64    `json:"province_id"`
			RegencyID   uint64    `json:"regency_id"`
			OpenedAt    time.Time `json:"opened_at" time_format"sql_date"`
		}
	}
)

func ValidateRequest(req *UpdateSchoolRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateSchoolRequest, accountID uint64) domain.School {
	acc := strconv.FormatUint(accountID, 10)
	return domain.School{
		Name:        req.Data.Name,
		Phone:       req.Data.Phone,
		Email:       req.Data.Email,
		Address:     req.Data.Address,
		PosCode:     req.Data.PosCode,
		Description: req.Data.Description,
		ProvinceID:  req.Data.ProvinceID,
		RegencyID:   req.Data.RegencyID,
		OpenedAt:    req.Data.OpenedAt,
		UpdatedBy:   acc,
	}
}
