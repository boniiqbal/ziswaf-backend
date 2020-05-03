package list_prognosis

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListPrognosisResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListPrognosisResponseData `json:"data"`
	}

	ListPrognosisResponseData struct {
		ID           uint64    `json:"id"`
		Total        uint64    `json:"total"`
		Month        int       `json:"month"`
		Year         int       `json:"year"`
		DivisionID   uint64    `json:"division_id"`
		DivisionName string    `json:"division_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func (res *ListPrognosisResponse) AddDomain(prognosis []domain.TransactionGoal, division []domain.Division) {
	response := ListPrognosisResponseData{}

	for _, ctg := range prognosis {
		for _, reg := range division {
			if ctg.DivisionID == reg.ID {
				response.DivisionName = reg.Name
				response.DivisionID = reg.ID
			}
			response.ID = ctg.ID
			response.Total = ctg.Total
			response.Month = ctg.Month
			response.Year = ctg.Year
			response.CreatedAt = ctg.CreatedAt
			response.UpdatedAt = ctg.UpdatedAt
		}
		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListPrognosisResponseData, pagination base.PaginationResponseData, message string, success bool) ListPrognosisResponse {
	return ListPrognosisResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.TransactionGoal, division []domain.Division) []ListPrognosisResponseData {
	response := ListPrognosisResponse{}

	response.AddDomain(domain, division)
	return response.Data
}
