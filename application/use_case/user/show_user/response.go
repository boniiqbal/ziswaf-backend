package show_user

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowUserResponse struct {
		base.BaseResponse
		Data ShowUserResponseData `json:"data"`
	}

	ShowUserResponseData struct {
		ID        uint64           `json:"id"`
		Username  string           `json:"username"`
		CreatedBy string           `json:"created_by"`
		UpdatedBy string           `json:"updated_by"`
		Status    int              `json:"status"`
		Role      int              `json:"role"`
		LastLogin time.Time        `json:"last_login"`
		School    SchoolResponse   `json:"school"`
		Employee  EmployeeResponse `json:"employee"`
		CreatedAt time.Time        `json:"created_at"`
		UpdatedAt time.Time        `json:"updated_at"`
	}

	SchoolResponse struct {
		ID      uint64 `json:"id"`
		Name    string `json:"name"`
		PosCode int    `json:"pos_code"`
	}

	EmployeeResponse struct {
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func SetResponse(domain ShowUserResponseData, message string, success bool) ShowUserResponse {
	return ShowUserResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.User) ShowUserResponseData {
	var (
		schResponse SchoolResponse
		empResponse EmployeeResponse
	)

	empResponse.ID = domain.Employee.ID
	empResponse.Email = domain.Employee.Email
	empResponse.Name = domain.Employee.Name
	schResponse.ID = domain.Employee.School.ID
	schResponse.Name = domain.Employee.School.Name
	schResponse.PosCode = domain.Employee.School.PosCode

	return ShowUserResponseData{
		ID:        domain.ModelSoftDelete.ID,
		Username:  domain.Username,
		CreatedBy: domain.CreatedBy,
		UpdatedBy: domain.UpdatedBy,
		Status:    domain.Status,
		Role:      domain.Role,
		LastLogin: domain.LastLogin,
		School:    schResponse,
		Employee:  empResponse,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
