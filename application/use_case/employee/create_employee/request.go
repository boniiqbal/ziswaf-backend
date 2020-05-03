package create_employee

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateEmployeeRequest struct {
		SchoolID       uint64    `json:"school_id" form:"school_id"`
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
)

func ValidateRequest(req *CreateEmployeeRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateEmployeeRequest, schoolID uint64) domain.Employee {
	return domain.Employee{
		SchoolID:       schoolID,
		Name:           req.Name,
		PlaceOfBirth:   req.PlaceOfBirth,
		BirthOfDate:    req.BirthOfDate,
		Phone:          req.Phone,
		Email:          req.Email,
		Address:        req.Address,
		Status:         req.Status,
		RegisteredYear: req.RegisteredYear,
		PosCode:        req.PosCode,
		ProvinceID:     req.ProvinceID,
		RegencyID:      req.RegencyID,
	}
}
