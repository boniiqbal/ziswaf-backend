package update_school

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateSchoolResponse struct {
		base.BaseResponse
		Data UpdateSchoolResponseData `json:"data"`
	}

	UpdateSchoolResponseData struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name" validate:"required"`
		Phone       string    `json:"phone" validate:"numeric"`
		Email       string    `json:"email" validate:"email"`
		Address     string    `json:"address"`
		PosCode     int       `json:"pos_code"`
		Description string    `json:"description"`
		ProvinceID  uint64    `json:"province_id"`
		RegencyID   uint64    `json:"regency_id"`
		OpenedAt    time.Time `json:"opened_at" time_format"sql_date"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateSchoolResponseData, message string, success bool) UpdateSchoolResponse {
	return UpdateSchoolResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.School) UpdateSchoolResponseData {
	return UpdateSchoolResponseData{
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
		CreatedAt:   domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:   domain.ModelSoftDelete.UpdatedAt,
	}
}
