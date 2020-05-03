package delete_user

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteUserResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteUserResponse {
	return DeleteUserResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
