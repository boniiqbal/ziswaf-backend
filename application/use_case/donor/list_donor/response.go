package list_donor

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListDonorsResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListDonorsResponseData    `json:"data"`
	}

	ListDonorsResponseData struct {
		ID           uint64    `json:"id"`
		Name         string    `json:"name"`
		CompanyName  string    `json:"company_name"`
		IsCompany    bool      `json:"is_company"`
		Position     string    `json:"position"`
		Email        string    `json:"email"`
		Address      string    `json:"address"`
		Phone        string    `json:"phone"`
		Status       int       `json:"status"`
		Npwp         int64     `json:"npwp"`
		PosCode      int       `json:"pos_code"`
		Info         string    `json:"info"`
		ProvinceName string    `json:"province_name"`
		RegencyName  string    `json:"regency_name"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func (res *ListDonorsResponse) AddDomain(donors []domain.Donor, regency []domain.Regency) {
	response := ListDonorsResponseData{}

	for _, ctg := range donors {
		for _, reg := range regency {
			if ctg.RegencyID == reg.ID {
				response.RegencyName = reg.Name
				response.ProvinceName = reg.Province.Name
			}
			response.ID = ctg.ModelSoftDelete.ID
			response.Name = ctg.Name
			response.CompanyName = ctg.CompanyName
			response.IsCompany = ctg.IsCompany
			response.Position = ctg.Position
			response.Email = ctg.Email
			response.Address = ctg.Address
			response.Phone = ctg.Phone
			response.Status = ctg.Status
			response.Npwp = ctg.Npwp
			response.PosCode = ctg.PosCode
			response.Info = ctg.Info
			response.CreatedAt = ctg.ModelSoftDelete.CreatedAt
			response.UpdatedAt = ctg.ModelSoftDelete.UpdatedAt
		}
		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListDonorsResponseData, pagination base.PaginationResponseData, message string, success bool) ListDonorsResponse {
	return ListDonorsResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.Donor, regency []domain.Regency) []ListDonorsResponseData {
	response := ListDonorsResponse{}

	response.AddDomain(domain, regency)
	return response.Data
}
