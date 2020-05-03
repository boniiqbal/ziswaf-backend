package update_category

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateCategoryResponse struct {
		base.BaseResponse
		Data UpdateCategoryResponseData `json:"data"`
	}

	UpdateCategoryResponseData struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Status      int       `json:"status"`
		CreatedBy   string    `json:"created_by"`
		UpdatedBy   string    `json:"updated_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateCategoryResponseData, message string, success bool) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Category) UpdateCategoryResponseData {
	return UpdateCategoryResponseData{
		Name:        domain.Name,
		Description: domain.Description,
		Status:      domain.Status,
		CreatedBy:   domain.CreatedBy,
		UpdatedBy:   domain.UpdatedBy,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}