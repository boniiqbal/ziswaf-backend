package list_province

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListProvincesResponse struct {
		base.BaseResponse
		Data []ListProvincesResponseData `json:"data"`
	}

	ListProvincesResponseData struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (res *ListProvincesResponse) AddDomain(provinces []domain.Province) {
	response := ListProvincesResponseData{}

	for _, prov := range provinces {
		response.ID = prov.Model.ID
		response.Name = prov.Name
		response.CreatedAt = prov.Model.CreatedAt
		response.UpdatedAt = prov.Model.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListProvincesResponseData, message string, success bool) ListProvincesResponse {

	return ListProvincesResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.Province) []ListProvincesResponseData {
	response := ListProvincesResponse{}

	response.AddDomain(domain)
	return response.Data
}
