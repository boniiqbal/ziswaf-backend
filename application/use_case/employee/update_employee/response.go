package update_employee

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateEmployeeResponse struct {
		base.BaseResponse
		Data UpdateEmployeeResponseData `json:"data"`
	}

	UpdateEmployeeResponseData struct {
		ID             uint64    `json:"id"`
		Name           string    `json:"name"`
		PlaceOfBirth   string    `json:"place_of_birth"`
		BirthOfDate    time.Time `json:"birth_of_date"`
		Phone          string    `json:"phone"`
		Email          string    `json:"email"`
		Address        string    `json:"address"`
		Status         int       `json:"status"`
		RegisteredYear time.Time `json:"registered_year"`
		PosCode        int       `json:"pos_code"`
		ProvinceID     uint64    `json:"province_id"`
		RegencyID      uint64    `json:"regency_id"`
		Image          string    `json:"image"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateEmployeeResponseData, message string, success bool) UpdateEmployeeResponse {
	return UpdateEmployeeResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Employee) UpdateEmployeeResponseData {
	return UpdateEmployeeResponseData{
		ID:             domain.ModelSoftDelete.ID,
		Name:           domain.Name,
		PlaceOfBirth:   domain.PlaceOfBirth,
		BirthOfDate:    domain.BirthOfDate,
		Phone:          domain.Phone,
		Email:          domain.Email,
		Address:        domain.Address,
		Status:         domain.Status,
		RegisteredYear: domain.RegisteredYear,
		PosCode:        domain.PosCode,
		ProvinceID:     domain.ProvinceID,
		RegencyID:      domain.RegencyID,
		Image:          domain.Image,
		CreatedAt:      domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:      domain.ModelSoftDelete.UpdatedAt,
	}
}
