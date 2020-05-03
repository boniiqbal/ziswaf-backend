package record_school

import (
	base "github.com/refactory-id/go-core-package/response"
	domain "ziswaf-backend/domain/entities"
)

type (
	RecordSchoolResponse struct {
		base.BaseResponse
		Data RecordSchoolResponseData `json:"data"`
	}

	RecordSchoolResponseData struct {
		ID             uint64 `json:"id"`
		Name           string `json:"name"`
		DonationRecord int    `json:"donation_record"`
		EmployeeRecord int    `json:"personel_record"`
		StudentRecord  int    `json:"student_record"`
	}
)

func SetResponse(domain RecordSchoolResponseData, message string, success bool) RecordSchoolResponse {
	return RecordSchoolResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.ReportSchool) RecordSchoolResponseData {
	return RecordSchoolResponseData{
		ID:             domain.ID,
		Name:           domain.Name,
		DonationRecord: domain.TotalDonation,
		EmployeeRecord: domain.TotalEmployee,
		StudentRecord:  domain.TotalStudent,
	}
}
