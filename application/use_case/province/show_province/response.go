package show_province

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowProvinceResponse struct {
		base.BaseResponse
		Data ShowProvinceResponseData `json:"data"`
	}

	ShowProvinceResponseData struct {
		ID        uint64            `json:"id"`
		Name      string            `json:"name"`
		Regency   []RegencyResponse `json:"regency"`
		CreatedAt time.Time         `json:"created_at"`
		UpdatedAt time.Time         `json:"updated_at"`
	}

	RegencyResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func SetResponse(domain ShowProvinceResponseData, message string, success bool) ShowProvinceResponse {
	return ShowProvinceResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Province) ShowProvinceResponseData {
	var (
		regResponse RegencyResponse
		response    ShowProvinceResponseData
	)

	for _, v := range domain.Regency {
		regResponse.ID = v.ID
		regResponse.Name = v.Name

		response.Regency = append(response.Regency, regResponse)
	}

	return ShowProvinceResponseData{
		ID:        domain.Model.ID,
		Name:      domain.Name,
		Regency:   response.Regency,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
