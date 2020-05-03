package delete_donor

import (
	base "github.com/refactory-id/go-core-package/response"
)

type (
	DeleteDonorResponse struct {
		base.BaseResponse
	}
)

func SetResponse(message string, success bool) DeleteDonorResponse {
	return DeleteDonorResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
