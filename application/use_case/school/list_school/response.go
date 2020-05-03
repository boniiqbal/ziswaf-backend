package list_school

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListSchoolResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListSchoolResponseData    `json:"data"`
	}

	ListSchoolResponseData struct {
		ID           uint64               `json:"id"`
		Name         string               `json:"name"`
		Phone        string               `json:"phone"`
		Email        string               `json:"email"`
		Address      string               `json:"address"`
		PosCode      int                  `json:"pos_code"`
		Description  string               `json:"description"`
		CreatedBy    string               `json:"created_by"`
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

func (res *ListSchoolResponse) AddDomain(school []domain.School) {
	response := ListSchoolResponseData{}

	for _, reg := range school {
		response.ID = reg.ModelSoftDelete.ID
		response.Name = reg.Name
		response.Phone = reg.Phone
		response.Email = reg.Email
		response.Address = reg.Address
		response.PosCode = reg.PosCode
		response.Description = reg.Description
		response.ProvinceName = reg.Province.Name
		response.RegencyName = reg.Regency.Name
		response.CreatedBy = reg.CreatedBy
		response.OpenedAt = reg.OpenedAt
		response.CreatedAt = reg.ModelSoftDelete.CreatedAt
		response.UpdatedAt = reg.ModelSoftDelete.UpdatedAt
		response.HeadMaster = HeadMaster{
			ID:    reg.HeadMaster.ID,
			Name:  reg.HeadMaster.Name,
			Email: reg.HeadMaster.Email,
			Phone: reg.HeadMaster.Phone,
			Image: reg.HeadMaster.Image,
		}
		response.Report = ReportSchoolResponse{
			TotalEmployee:       reg.Report.TotalEmployee,
			TotalSosialEmployee: reg.Report.TotalSosialEmployee,
			TotalStudent:        reg.Report.TotalStudent,
			TotalTeacher:        reg.Report.TotalTeacher,
		}

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListSchoolResponseData, pagination base.PaginationResponseData, message string, success bool) ListSchoolResponse {

	return ListSchoolResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.School) []ListSchoolResponseData {
	response := ListSchoolResponse{}

	response.AddDomain(domain)
	return response.Data
}
