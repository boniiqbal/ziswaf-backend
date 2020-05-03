package show_employee

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowEmployeeResponse struct {
		base.BaseResponse
		Data ShowEmployeeResponseData `json:"data"`
	}

	ShowEmployeeResponseData struct {
		ID             uint64    `json:"id"`
		SchoolName     string    `json:"school_name"`
		Name           string    `json:"name"`
		PlaceOfBirth   string    `json:"place_of_birth"`
		BirthOfDate    time.Time `json:"birth_of_date"`
		Phone          string    `json:"phone"`
		Email          string    `json:"email"`
		Address        string    `json:"address"`
		Status         int       `json:"status"`
		RegisteredYear time.Time `json:"registered_year"`
		PosCode        int       `json:"pos_code"`
		ProvinceName   string    `json:"province_name"`
		RegencyName    string    `json:"regency_name"`
		Image          string    `json:"image"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

func SetResponse(domain ShowEmployeeResponseData, message string, success bool) ShowEmployeeResponse {
	return ShowEmployeeResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Employee) ShowEmployeeResponseData {
	return ShowEmployeeResponseData{
		ID:             domain.ModelSoftDelete.ID,
		SchoolName:     domain.School.Name,
		Name:           domain.Name,
		PlaceOfBirth:   domain.PlaceOfBirth,
		BirthOfDate:    domain.BirthOfDate,
		Phone:          domain.Phone,
		Email:          domain.Email,
		Address:        domain.Address,
		Status:         domain.Status,
		RegisteredYear: domain.RegisteredYear,
		PosCode:        domain.PosCode,
		ProvinceName:   domain.Province.Name,
		RegencyName:    domain.Regency.Name,
		Image:          domain.Image,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
