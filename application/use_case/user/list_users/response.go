package list_users

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListUsersResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListUsersResponseData     `json:"data"`
	}

	ListUsersResponseData struct {
		ID        uint64           `json:"id"`
		Username  string           `json:"username"`
		CreatedBy string           `json:"created_by"`
		UpdatedBy string           `json:"updated_by"`
		Status    int              `json:"status"`
		Role      int              `json:"role"`
		School    SchoolResponse   `json:"school"`
		Employee  EmployeeResponse `json:"employee"`
		LastLogin time.Time        `json:"last_login"`
		CreatedAt time.Time        `json:"created_at"`
		UpdatedAt time.Time        `json:"updated_at"`
	}

	SchoolResponse struct {
		ID      uint64 `json:"id"`
		Name    string `json:"name"`
		PosCode int    `json:"pos_code"`
	}

	EmployeeResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func (res *ListUsersResponse) AddDomain(users []domain.User) {
	response := ListUsersResponseData{}

	for _, prd := range users {
		response.Employee.Email = prd.Employee.Email
		response.Employee.Name = prd.Employee.Name
		response.School.ID = prd.Employee.School.ID
		response.School.Name = prd.Employee.School.Name
		response.School.PosCode = prd.Employee.School.PosCode
		response.ID = prd.ModelSoftDelete.ID
		response.Username = prd.Username
		response.CreatedBy = prd.CreatedBy
		response.UpdatedBy = prd.UpdatedBy
		response.Status = prd.Status
		response.Role = prd.Role
		response.LastLogin = prd.LastLogin
		response.CreatedAt = prd.CreatedAt
		response.UpdatedAt = prd.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListUsersResponseData, pagination base.PaginationResponseData, message string, success bool) ListUsersResponse {

	return ListUsersResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.User) []ListUsersResponseData {
	response := ListUsersResponse{}

	response.AddDomain(domain)
	return response.Data
}
