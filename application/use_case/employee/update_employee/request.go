package update_employee

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateEmployeeRequest struct {
		Data struct {
			Name           string    `json:"name" form:"name" validate:"required"`
			PlaceOfBirth   string    `json:"place_of_birth" form:"place_of_birth"`
			BirthOfDate    time.Time `json:"birth_of_date" form:"birth_of_date"`
			Phone          string    `json:"phone" form:"phone"`
			Email          string    `json:"email" form:"email"`
			Address        string    `json:"address" form:"address"`
			Status         int       `json:"status" form:"status"`
			RegisteredYear time.Time `json:"registered_year" form:"registered_year"`
			PosCode        int       `json:"pos_code" form:"pos_code"`
			ProvinceID     uint64    `json:"province_id" form:"province_id"`
			RegencyID      uint64    `json:"regency_id" form:"regency_id"`
		}
	}
)

func ValidateRequest(req *UpdateEmployeeRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateEmployeeRequest) domain.Employee {
	return domain.Employee{
		Name:           req.Data.Name,
		PlaceOfBirth:   req.Data.PlaceOfBirth,
		BirthOfDate:    req.Data.BirthOfDate,
		Phone:          req.Data.Phone,
		Email:          req.Data.Email,
		Address:        req.Data.Address,
		Status:         req.Data.Status,
		RegisteredYear: req.Data.RegisteredYear,
		PosCode:        req.Data.PosCode,
		ProvinceID:     req.Data.ProvinceID,
		RegencyID:      req.Data.RegencyID,
	}
}
