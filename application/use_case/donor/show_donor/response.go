package show_donor

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowDonorResponse struct {
		base.BaseResponse
		Data ShowDonorResponseData `json:"data"`
	}

	ShowDonorResponseData struct {
		ID           uint64    `json:"id"`
		Name         string    `json:"name"`
		CompanyName  string    `json:"company_name"`
		IsCompany    bool      `json:"is_company"`
		Position     string    `json:"position"`
		Email        string    `json:"email"`
		Address      string    `json:"address"`
		Phone        string    `json:"phone"`
		Status       int       `json:"status"`
		Npwp         int64     `json:"npwp"`
		PosCode      int       `json:"pos_code"`
		Info         string    `json:"info"`
		ProvinceName string    `json:"province_name"`
		RegencyName  string    `json:"regency_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func SetResponse(domain ShowDonorResponseData, message string, success bool) ShowDonorResponse {
	return ShowDonorResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Donor, regency []domain.Regency) ShowDonorResponseData {
	var (
		regName  string
		provName string
	)

	for _, regency := range regency {
		if domain.RegencyID == regency.ID {
			regName = regency.Name
			provName = regency.Province.Name
		}
	}
	return ShowDonorResponseData{
		ID:           domain.ModelSoftDelete.ID,
		Name:         domain.Name,
		CompanyName:  domain.CompanyName,
		IsCompany:    domain.IsCompany,
		Position:     domain.Position,
		Email:        domain.Email,
		Address:      domain.Address,
		Phone:        domain.Phone,
		Status:       domain.Status,
		Npwp:         domain.Npwp,
		PosCode:      domain.PosCode,
		Info:         domain.Info,
		ProvinceName: provName,
		RegencyName:  regName,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
