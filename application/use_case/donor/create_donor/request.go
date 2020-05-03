package create_donor

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreateDonorRequest struct {
		Data struct {
			Name        string `json:"name" validate:"required"`
			CompanyName string `json:"company_name"`
			IsCompany   bool   `json:"is_company"`
			Position    string `json:"position"`
			Email       string `json:"email" validate:"required,email"`
			Address     string `json:"address"`
			Phone       string `json:"phone" validate:"required"`
			Status      int    `json:"status"`
			Npwp        int64  `json:"npwp"`
			PosCode     int    `json:"pos_code"`
			Info        string `json:"info"`
			ProvinceID  uint64 `json:"province_id" validate:"required"`
			RegencyID   uint64 `json:"regency_id" validate:"required"`
		}
	}
)

func ValidateRequest(req *CreateDonorRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateDonorRequest) domain.Donor {
	return domain.Donor{
		Name:        req.Data.Name,
		CompanyName: req.Data.CompanyName,
		IsCompany:   req.Data.IsCompany,
		Position:    req.Data.Position,
		Email:       req.Data.Email,
		Address:     req.Data.Address,
		Phone:       req.Data.Phone,
		Status:      req.Data.Status,
		Npwp:        req.Data.Npwp,
		PosCode:     req.Data.PosCode,
		Info:        req.Data.Info,
		ProvinceID:  req.Data.ProvinceID,
		RegencyID:   req.Data.RegencyID,
	}
}
