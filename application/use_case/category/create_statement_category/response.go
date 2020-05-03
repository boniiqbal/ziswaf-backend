package create_statement_category

import (
	base "github.com/refactory-id/go-core-package/response"
	domain "ziswaf-backend/domain/entities"
)

type (
	CreateStatementCategoryResponse struct {
		base.BaseResponse
		Data CreateStatementCategoryResponseData `json:"data"`
	}

	CreateStatementCategoryResponseData struct {
		ID   uint64 `json:"id"`
		Name string `json:"jenis_kategory"`
	}
)

func SetResponse(domain CreateStatementCategoryResponseData, message string, success bool) CreateStatementCategoryResponse {
	return CreateStatementCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.StatementCategory) CreateStatementCategoryResponseData {
	return CreateStatementCategoryResponseData{
		ID:   domain.ModelSoftDelete.ID,
		Name: domain.Name,
	}
}
