package create_user

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateUserResponse struct {
		base.BaseResponse
		Data CreateUserResponseData `json:"data"`
	}

	CreateUserResponseData struct {
		ID        uint64    `json:"id"`
		Username  string    `json:"username"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		Status    int       `json:"status"`
		Role      int       `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func SetResponse(domain CreateUserResponseData, message string, success bool) CreateUserResponse {
	return CreateUserResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.User) CreateUserResponseData {
	return CreateUserResponseData{
		ID:        domain.ModelSoftDelete.ID,
		Username:  domain.Username,
		CreatedBy: domain.CreatedBy,
		UpdatedBy: domain.UpdatedBy,
		Status:    domain.Status,
		Role:      domain.Role,
		CreatedAt: domain.ModelSoftDelete.CreatedAt,
		UpdatedAt: domain.ModelSoftDelete.UpdatedAt,
	}
}
