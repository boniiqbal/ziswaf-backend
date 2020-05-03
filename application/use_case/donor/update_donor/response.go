package update_donor

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateDonorResponse struct {
		base.BaseResponse
		Data UpdateDonorResponseData `json:"data"`
	}

	UpdateDonorResponseData struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		CompanyName string    `json:"company_name"`
		IsCompany   bool      `json:"is_company"`
		Position    string    `json:"position"`
		Email       string    `json:"email"`
		Address     string    `json:"address"`
		Phone       string    `json:"phone"`
		Status      int       `json:"status"`
		Npwp        int64     `json:"npwp"`
		PosCode     int       `json:"pos_code"`
		Info        string    `json:"info"`
		ProvinceID  uint64    `json:"province_id"`
		RegencyID   uint64    `json:"regency_id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateDonorResponseData, message string, success bool) UpdateDonorResponse {
	return UpdateDonorResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Donor) UpdateDonorResponseData {
	return UpdateDonorResponseData{
		ID:          domain.ModelSoftDelete.ID,
		Name:        domain.Name,
		CompanyName: domain.CompanyName,
		IsCompany:   domain.IsCompany,
		Position:    domain.Position,
		Email:       domain.Email,
		Address:     domain.Address,
		Phone:       domain.Phone,
		Status:      domain.Status,
		Npwp:        domain.Npwp,
		PosCode:     domain.PosCode,
		Info:        domain.Info,
		ProvinceID:  domain.ProvinceID,
		RegencyID:   domain.RegencyID,
		CreatedAt:   domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:   domain.ModelSoftDelete.UpdatedAt,
	}
}
