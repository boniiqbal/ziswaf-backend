package list_category

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListCategoriesResponse struct {
		base.BaseResponse
		Data []ListCategoriesResponseData `json:"data"`
	}

	ListCategoriesResponseData struct {
		ID                  uint64      `json:"id"`
		Name                string      `json:"name"`
		Description         string      `json:"description"`
		Status              int         `json:"status"`
		CreatedBy           string      `json:"created_by"`
		UpdatedBy           string      `json:"updated_by"`
		CreatedAt           time.Time   `json:"created_at"`
		UpdatedAt           time.Time   `json:"updated_at"`
		StatementCategories interface{} `json:"statement_categories"`
	}
)

func (res *ListCategoriesResponse) AddDomain(categories []domain.Category) {
	response := ListCategoriesResponseData{}

	for _, ctg := range categories {

		response.ID = ctg.Model.ID
		response.Name = ctg.Name
		response.Description = ctg.Description
		response.CreatedBy = ctg.CreatedBy
		response.UpdatedBy = ctg.UpdatedBy
		response.Status = ctg.Status
		response.CreatedAt = ctg.Model.CreatedAt
		response.UpdatedAt = ctg.Model.UpdatedAt
		response.StatementCategories = ctg.StatementCategories

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListCategoriesResponseData, message string, success bool) ListCategoriesResponse {

	return ListCategoriesResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.Category) []ListCategoriesResponseData {
	response := ListCategoriesResponse{}

	response.AddDomain(domain)
	return response.Data
}
