package list_village

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListVillageResponse struct {
		base.BaseResponse
		Data []ListVillageResponseData `json:"data"`
	}

	ListVillageResponseData struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func (res *ListVillageResponse) AddDomain(village []domain.Village) {
	response := ListVillageResponseData{}

	for _, reg := range village {
		response.ID = reg.Model.ID
		response.Name = reg.Name

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListVillageResponseData, message string, success bool) ListVillageResponse {

	return ListVillageResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.Village) []ListVillageResponseData {
	response := ListVillageResponse{}

	response.AddDomain(domain)
	return response.Data
}
