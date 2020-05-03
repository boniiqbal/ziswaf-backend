package update_prognosis

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdatePrognosisResponse struct {
		base.BaseResponse
		Data UpdatePrognosisResponseData `json:"data"`
	}

	UpdatePrognosisResponseData struct {
		ID         uint64    `json:"id"`
		Total      uint64    `json:"total"`
		Month      int       `json:"month"`
		Year       int       `json:"year"`
		DivisionID uint64    `json:"division_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdatePrognosisResponseData, message string, success bool) UpdatePrognosisResponse {
	return UpdatePrognosisResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.TransactionGoal) UpdatePrognosisResponseData {
	return UpdatePrognosisResponseData{
		ID:         domain.Model.ID,
		Total:      domain.Total,
		Month:      domain.Month,
		Year:       domain.Year,
		DivisionID: domain.DivisionID,
		CreatedAt:  domain.Model.CreatedAt,
		UpdatedAt:  domain.Model.UpdatedAt,
	}
}
