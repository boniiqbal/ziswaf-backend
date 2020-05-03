package update_status

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateStatusResponse struct {
		base.BaseResponse
		Data UpdateStatusResponseData `json:"data"`
	}

	UpdateStatusResponseData struct {
		ID        uint64    `json:"id"`
		Name      string    `json:"name"`
		Username  string    `json:"username"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		Status    int       `json:"status"`
		Role      int       `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateStatusResponseData, message string, success bool) UpdateStatusResponse {
	return UpdateStatusResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.User) UpdateStatusResponseData {
	return UpdateStatusResponseData{
		ID:        domain.ModelSoftDelete.ID,
		Name:      domain.Name,
		Username:  domain.Username,
		CreatedBy: domain.CreatedBy,
		UpdatedBy: domain.UpdatedBy,
		Status:    domain.Status,
		Role:      domain.Role,
		CreatedAt: domain.ModelSoftDelete.CreatedAt,
		UpdatedAt: domain.ModelSoftDelete.UpdatedAt,
	}
}
