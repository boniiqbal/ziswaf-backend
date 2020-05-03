package list_regency

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListRegenciesResponse struct {
		base.BaseResponse
		Data []ListRegenciesResponseData `json:"data"`
	}

	ListRegenciesResponseData struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		ProvinceName string `json:"province_name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (res *ListRegenciesResponse) AddDomain(regencies []domain.Regency) {
	response := ListRegenciesResponseData{}

	for _, reg := range regencies {
		response.ID = reg.Model.ID
		response.Name = reg.Name
		response.ProvinceName = reg.Province.Name
		response.CreatedAt = reg.Model.CreatedAt
		response.UpdatedAt = reg.Model.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListRegenciesResponseData, message string, success bool) ListRegenciesResponse {

	return ListRegenciesResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.Regency) []ListRegenciesResponseData {
	response := ListRegenciesResponse{}

	response.AddDomain(domain)
	return response.Data
}
