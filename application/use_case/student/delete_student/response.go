package delete_student

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteStudentResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteStudentResponse {
	return DeleteStudentResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
