package create_prognosis

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreatePrognosisResponse struct {
		base.BaseResponse
		Data []CreatePrognosisResponseData `json:"data"`
	}

	CreatePrognosisResponseData struct {
		Total      uint64    `json:"total"`
		Month      int       `json:"month"`
		Year       int       `json:"year"`
		DivisionID uint64    `json:"division_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func (res *CreatePrognosisResponse) AddDomain(prognosis []domain.TransactionGoal) {
	response := CreatePrognosisResponseData{}

	for _, prog := range prognosis {

		response.Total = prog.Total
		response.Month = prog.Month
		response.Year = prog.Year
		response.DivisionID = prog.DivisionID
		response.CreatedAt = time.Now()
		response.UpdatedAt = time.Now()

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []CreatePrognosisResponseData, message string, success bool) CreatePrognosisResponse {
	return CreatePrognosisResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.TransactionGoal) []CreatePrognosisResponseData {
	response := CreatePrognosisResponse{}

	response.AddDomain(domain)
	return response.Data
}
