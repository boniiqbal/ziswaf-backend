package show_category

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowCategoryResponse struct {
		base.BaseResponse
		Data ShowCategoryResponseData `json:"data"`
	}

	ShowCategoryResponseData struct {
		ID           uint64    `json:"id"`
		Name         string    `json:"name"`
		DivisionName string    `json:"division_name"`
		Description  string    `json:"description"`
		Status       int       `json:"status"`
		CreatedBy    string    `json:"created_by"`
		UpdatedBy    string    `json:"updated_by"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func SetResponse(domain ShowCategoryResponseData, message string, success bool) ShowCategoryResponse {
	return ShowCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Category) ShowCategoryResponseData {
	return ShowCategoryResponseData{
		ID:           domain.Model.ID,
		Name:         domain.Name,
		DivisionName: domain.Name,
		Description:  domain.Description,
		Status:       domain.Status,
		CreatedBy:    domain.CreatedBy,
		UpdatedBy:    domain.UpdatedBy,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
