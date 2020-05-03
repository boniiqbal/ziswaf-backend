package show_division

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowDivisionResponse struct {
		base.BaseResponse
		Data ShowDivisionResponseData `json:"data"`
	}

	ShowDivisionResponseData struct {
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

func SetResponse(domain ShowDivisionResponseData, message string, success bool) ShowDivisionResponse {
	return ShowDivisionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Division) ShowDivisionResponseData {
	return ShowDivisionResponseData{
		ID:          domain.Model.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Status:      domain.Status,
		CreatedBy:   domain.CreatedBy,
		UpdatedBy:   domain.UpdatedBy,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
