package delete_employee

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteEmployeeResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteEmployeeResponse {
	return DeleteEmployeeResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
