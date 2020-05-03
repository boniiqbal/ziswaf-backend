package change_password

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ChangePasswordResponse struct {
		base.BaseResponse
		Data ChangePasswordResponseData `json:"data"`
	}

	ChangePasswordResponseData struct {
		ID        uint64         `json:"id"`
		Username  string         `json:"username"`
		Name      string         `json:"name"`
		Status    int            `json:"status"`
		Role      int            `json:"role"`
		School    SchoolResponse `json:"school,omitempty"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
	}

	SchoolResponse struct {
		ID      uint64 `json:"id"`
		Name    string `json:"name"`
		PosCode int    `json:"pos_code"`
	}
)

func SetResponse(domain ChangePasswordResponseData, message string, success bool) ChangePasswordResponse {
	return ChangePasswordResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.User) ChangePasswordResponseData {
	var schResponse SchoolResponse

	schResponse.ID = domain.Employee.School.ID
	schResponse.Name = domain.Employee.School.Name
	schResponse.PosCode = domain.Employee.School.PosCode

	return ChangePasswordResponseData{
		ID:        domain.ID,
		Name:      domain.Name,
		Username:  domain.Username,
		Status:    domain.Status,
		Role:      domain.Role,
		School:    schResponse,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
