package delete_statement_category

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteStatementCategoryResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteStatementCategoryResponse {
	return DeleteStatementCategoryResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
