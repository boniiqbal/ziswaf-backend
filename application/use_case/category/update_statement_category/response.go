package update_statement_category

import (
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateStatementCategoryResponse struct {
		base.BaseResponse
		Data UpdateStatementCategoryResponseData `json:"data"`
	}

	UpdateStatementCategoryResponseData struct {
		Name       string `json:"name"`
		CategoryID uint64 `json:"category_id"`
	}
)

func SetResponse(domain UpdateStatementCategoryResponseData, message string, success bool) UpdateStatementCategoryResponse {
	return UpdateStatementCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.StatementCategory) UpdateStatementCategoryResponseData {
	return UpdateStatementCategoryResponseData{
		Name:       domain.Name,
		CategoryID: domain.CategoryID,
	}
}
