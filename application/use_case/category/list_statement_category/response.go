package list_statement_category

import (
	base "github.com/refactory-id/go-core-package/response"
	domain "ziswaf-backend/domain/entities"
)

type (
	ListStatementCategoriesResponse struct {
		base.BaseResponse
		Data []ListStatementCategoriesResponseData `json:"data"`
	}

	ListStatementCategoriesResponseData struct {
		ID   uint64 `json:"id"`
		Name string `json:"statement_category"`
	}
)

func (res *ListStatementCategoriesResponse) AddDomain(stCategories []domain.StatementCategory) {
	response := ListStatementCategoriesResponseData{}

	for _, ctg := range stCategories {
		response.ID = ctg.ModelSoftDelete.ID
		response.Name = ctg.Name

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListStatementCategoriesResponseData, message string, success bool) ListStatementCategoriesResponse {

	return ListStatementCategoriesResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Data: domain,
	}
}

func ResponseMapper(domain []domain.StatementCategory) []ListStatementCategoriesResponseData {
	response := ListStatementCategoriesResponse{}

	response.AddDomain(domain)
	return response.Data
}
