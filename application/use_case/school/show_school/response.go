package show_school

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ShowSchoolResponse struct {
		base.BaseResponse
		Data ShowSchoolResponseData `json:"data"`
	}

	ShowSchoolResponseData struct {
		ID           uint64               `json:"id"`
		Name         string               `json:"name"`
		Phone        string               `json:"phone"`
		Email        string               `json:"email"`
		Address      string               `json:"address"`
		PosCode      int                  `json:"pos_code"`
		Description  string               `json:"description"`
		ProvinceName string               `json:"province_name"`
		RegencyName  string               `json:"regency_name"`
		HeadMaster   HeadMaster           `json:"head_master"`
		Report       ReportSchoolResponse `json:"report_school"`
		OpenedAt     time.Time            `json:"opened_at"`
		CreatedAt    time.Time            `json:"created_at"`
		UpdatedAt    time.Time            `json:"updated_at"`
	}

	ReportSchoolResponse struct {
		TotalEmployee       int `json:"total_employee"`
		TotalTeacher        int `json:"total_teacher"`
		TotalSosialEmployee int `json:"total_social_employee"`
		TotalStudent        int `json:"total_student"`
	}

	HeadMaster struct {
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		Image string `json:"image"`
	}
)

func SetResponse(domain ShowSchoolResponseData, message string, success bool) ShowSchoolResponse {
	return ShowSchoolResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.School) ShowSchoolResponseData {
	headMaster := HeadMaster{}
	report := ReportSchoolResponse{}

	report.TotalEmployee = domain.Report.TotalEmployee
	report.TotalSosialEmployee = domain.Report.TotalSosialEmployee
	report.TotalTeacher = domain.Report.TotalTeacher
	report.TotalStudent = domain.Report.TotalStudent

	headMaster.ID = domain.HeadMaster.ID
	headMaster.Name = domain.HeadMaster.Name
	headMaster.Email = domain.HeadMaster.Email
	headMaster.Phone = domain.HeadMaster.Phone
	headMaster.Image = domain.HeadMaster.Image

	return ShowSchoolResponseData{
		ID:           domain.ModelSoftDelete.ID,
		Name:         domain.Name,
		Phone:        domain.Phone,
		Email:        domain.Email,
		Address:      domain.Address,
		PosCode:      domain.PosCode,
		Description:  domain.Description,
		ProvinceName: domain.Province.Name,
		RegencyName:  domain.Regency.Name,
		HeadMaster:   headMaster,
		Report:       report,
		OpenedAt:     domain.OpenedAt,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
