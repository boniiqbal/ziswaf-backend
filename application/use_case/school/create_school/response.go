package create_school

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateSchoolResponse struct {
		base.BaseResponse
		Data CreateSchoolResponseData `json:"data"`
	}

	CreateSchoolResponseData struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Phone       string    `json:"phone"`
		Email       string    `json:"email"`
		Address     string    `json:"address"`
		PosCode     int       `json:"pos_code"`
		Description string    `json:"description"`
		ProvinceID  uint64    `json:"province_id"`
		RegencyID   uint64    `json:"regency_id"`
		CreatedBy   string    `json:"created_by"`
		UpdatedBy   string    `json:"updated_by"`
		OpenedAt    time.Time `json:"opened_at"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func SetResponse(domain CreateSchoolResponseData, message string, success bool) CreateSchoolResponse {
	return CreateSchoolResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.School) CreateSchoolResponseData {
	return CreateSchoolResponseData{
		ID:          domain.ModelSoftDelete.ID,
		Name:        domain.Name,
		Phone:       domain.Phone,
		Email:       domain.Email,
		Address:     domain.Address,
		PosCode:     domain.PosCode,
		Description: domain.Description,
		ProvinceID:  domain.ProvinceID,
		RegencyID:   domain.RegencyID,
		OpenedAt:    domain.OpenedAt,
		CreatedBy:   domain.CreatedBy,
		UpdatedBy:   domain.UpdatedBy,
		CreatedAt:   domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:   domain.ModelSoftDelete.UpdatedAt,
	}
}
