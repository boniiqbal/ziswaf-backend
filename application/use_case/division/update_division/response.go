package update_division

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateDivisionResponse struct {
		base.BaseResponse
		Data UpdateDivisionResponseData `json:"data"`
	}

	UpdateDivisionResponseData struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Status      int       `json:"status"`
		UpdatedBy   string    `json:"updated_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateDivisionResponseData, message string, success bool) UpdateDivisionResponse {
	return UpdateDivisionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Division) UpdateDivisionResponseData {
	return UpdateDivisionResponseData{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Status:      domain.Status,
		UpdatedBy:   domain.UpdatedBy,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
