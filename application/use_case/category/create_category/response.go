package create_category

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateCategoryResponse struct {
		base.BaseResponse
		Data CreateCategoryResponseData `json:"data"`
	}

	CreateCategoryResponseData struct {
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

func SetResponse(domain CreateCategoryResponseData, message string, success bool) CreateCategoryResponse {
	return CreateCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Category) CreateCategoryResponseData {
	return CreateCategoryResponseData{
		ID:          domain.Model.ID,
		Name:        domain.Name,
		Description: domain.Description,
		Status:      domain.Status,
		CreatedBy:   domain.CreatedBy,
		UpdatedBy:   domain.UpdatedBy,
		CreatedAt:   domain.Model.CreatedAt,
		UpdatedAt:   domain.Model.UpdatedAt,
	}
}
