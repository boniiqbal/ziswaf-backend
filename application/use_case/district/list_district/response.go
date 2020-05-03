package list_district

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListDistrictResponse struct {
		base.BaseResponse
		Data []ListDistrictResponseData `json:"data"`
	}

	ListDistrictResponseData struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (res *ListDistrictResponse) AddDomain(district []domain.District) {
	response := ListDistrictResponseData{}

	for _, reg := range district {
		response.ID = reg.Model.ID
		response.Name = reg.Name
		response.CreatedAt = reg.Model.CreatedAt
		response.UpdatedAt = reg.Model.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListDistrictResponseData, message string, success bool) ListDistrictResponse {

	return ListDistrictResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.District) []ListDistrictResponseData {
	response := ListDistrictResponse{}

	response.AddDomain(domain)
	return response.Data
}
