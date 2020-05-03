package list_employee

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListEmployeesResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListEmployeesResponseData `json:"data"`
	}

	ListEmployeesResponseData struct {
		ID             uint64    `json:"id"`
		SchoolName     string    `json:"school_name"`
		Name           string    `json:"name"`
		PlaceOfBirth   string    `json:"place_of_birth"`
		BirthOfDate    time.Time `json:"birth_of_date"`
		Phone          string    `json:"phone"`
		Email          string    `json:"email"`
		Address        string    `json:"address"`
		Status         int       `json:"status"`
		RegisteredYear time.Time    `json:"registered_year"`
		PosCode        int       `json:"pos_code"`
		ProvinceName   string    `json:"province_name"`
		RegencyName    string    `json:"regency_name"`
		Image          string    `json:"image"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

func (res *ListEmployeesResponse) AddDomain(employees []domain.Employee) {
	response := ListEmployeesResponseData{}

	for _, emp := range employees {

		response.ID = emp.ModelSoftDelete.ID
		response.SchoolName = emp.School.Name
		response.Name = emp.Name
		response.PlaceOfBirth = emp.PlaceOfBirth
		response.BirthOfDate = emp.BirthOfDate
		response.Phone = emp.Phone
		response.Email = emp.Email
		response.Address = emp.Address
		response.Status = emp.Status
		response.RegisteredYear = emp.RegisteredYear
		response.PosCode = emp.PosCode
		response.ProvinceName = emp.Province.Name
		response.RegencyName = emp.Regency.Name
		response.Image = emp.Image
		response.CreatedAt = emp.ModelSoftDelete.CreatedAt
		response.UpdatedAt = emp.ModelSoftDelete.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListEmployeesResponseData, pagination base.PaginationResponseData, message string, success bool) ListEmployeesResponse {
	return ListEmployeesResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.Employee) []ListEmployeesResponseData {
	response := ListEmployeesResponse{}

	response.AddDomain(domain)
	return response.Data
}
