package login

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	LoginResponse struct {
		base.BaseResponse
		Data LoginResponseData `json:"data"`
	}

	LoginResponseData struct {
		UserID    uint64           `json:"user_id"`
		Token     string           `json:"token"`
		Username  string           `json:"username"`
		Status    int              `json:"status"`
		Role      int              `json:"role"`
		ExpiredAt int64            `json:"expired_at"`
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

func SetResponse(domain LoginResponseData, message string, success bool) LoginResponse {
	return LoginResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.User) LoginResponseData {
	var schResponse SchoolResponse
	var empResponse EmployeeResponse

	empResponse.ID = domain.Employee.ID
	empResponse.Email = domain.Employee.Email
	empResponse.Name = domain.Employee.Name

	schResponse.ID = domain.Employee.School.ID
	schResponse.Name = domain.Employee.School.Name
	schResponse.PosCode = domain.Employee.School.PosCode

	return LoginResponseData{
		UserID:    domain.ID,
		Token:     domain.AccessToken.Token,
		Username:  domain.Username,
		Status:    domain.Status,
		Role:      domain.Role,
		ExpiredAt: domain.AccessToken.ExpiredAt,
		School:    schResponse,
		Employee:  empResponse,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
