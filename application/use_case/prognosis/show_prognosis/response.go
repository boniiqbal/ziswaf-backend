package show_prognosis

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowPrognosisResponse struct {
		base.BaseResponse
		Data ShowPrognosisResponseData `json:"data"`
	}

	ShowPrognosisResponseData struct {
		ID           uint64    `json:"id"`
		Total        uint64    `json:"total"`
		Month        int       `json:"month"`
		Year         int       `json:"year"`
		DivisionName string    `json:"division_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func SetResponse(domain ShowPrognosisResponseData, message string, success bool) ShowPrognosisResponse {
	return ShowPrognosisResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.TransactionGoal, division []domain.Division) ShowPrognosisResponseData {
	var divName string

	for _, division := range division {
		if domain.DivisionID == division.ID {
			divName = division.Name
		}
	}
	return ShowPrognosisResponseData{
		ID:           domain.Model.ID,
		Total:        domain.Total,
		Month:        domain.Month,
		Year:         domain.Year,
		DivisionName: divName,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
