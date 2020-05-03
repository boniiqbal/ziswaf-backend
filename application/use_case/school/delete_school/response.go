package delete_school

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteSchoolResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteSchoolResponse {
	return DeleteSchoolResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
