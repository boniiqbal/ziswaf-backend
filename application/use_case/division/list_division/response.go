package list_division

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListDivisionsResponse struct {
		base.BaseResponse
		Data []ListDivisionsResponseData `json:"data"`
	}

	ListDivisionsResponseData struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Status      int       `json:"status"`
		CreatedBy   string    `json:"created_by"`
		UpdatedBy   string    `json:"updated_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func (res *ListDivisionsResponse) AddDomain(divisions []domain.Division) {
	response := ListDivisionsResponseData{}

	for _, ctg := range divisions {
		response.ID = ctg.Model.ID
		response.Name = ctg.Name
		response.Description = ctg.Description
		response.CreatedBy = ctg.CreatedBy
		response.UpdatedBy = ctg.UpdatedBy
		response.Status = ctg.Status
		response.CreatedAt = ctg.Model.CreatedAt
		response.UpdatedAt = ctg.Model.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListDivisionsResponseData, message string, success bool) ListDivisionsResponse {

	return ListDivisionsResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.Division) []ListDivisionsResponseData {
	response := ListDivisionsResponse{}

	response.AddDomain(domain)
	return response.Data
}
