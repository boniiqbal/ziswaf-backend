package list_student

import (
	"time"

	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	ListStudentResponse struct {
		base.BaseResponse
		Pagination base.PaginationResponseData `json:"pagination"`
		Data       []ListStudentResponseData   `json:"data"`
	}

	ListStudentResponseData struct {
		ID                      uint64    `json:"id"`
		IdentityNumber          string    `json:"identity_number"`
		SchoolName              string    `json:"school_name"`
		Name                    string    `json:"name"`
		Age                     string    `json:"age"`
		PlaceOfBirth            string    `json:"place_of_birth"`
		BirthOfDate             time.Time `json:"birth_of_date"`
		ChildRow                string    `json:"child_row"`
		TotalSibling            string    `json:"total_sibling"`
		Address                 string    `json:"address"`
		SosialStatus            int       `json:"sosial_status"`
		PosCode                 int       `json:"pos_code"`
		ProvinceName            string    `json:"province_name"`
		RegencyName             string    `json:"regency_name"`
		DistrictName            string    `json:"district_name"`
		VillageName             string    `json:"village_name"`
		EducationStatus         int       `json:"education_status"`
		RegisteredAt            time.Time `json:"registered_at"`
		FinishedAt              time.Time `json:"finished_at"`
		PunishmentCount         int       `json:"punishment_count"`
		PunishmentStart         time.Time `json:"punishment_start"`
		PunishmentEnd           time.Time `json:"punishment_end"`
		JuzKuranDescription     string    `json:"juz_kuran_description"`
		ChapterKuranDescription string    `json:"chapter_kuran_description"`
		HadistDescription       string    `json:"hadist_description"`
		EducationDescription    string    `json:"education_description"`
		ParentStatus            int       `json:"parent_status"`
		FatherName              string    `json:"father_name"`
		PlaceOfBirthFather      string    `json:"place_of_birth_father"`
		BirthOfDateFather       time.Time `json:"birth_of_date_father"`
		FatherOccupation        string    `json:"father_occupation"`
		FatherPhone             string    `json:"father_phone"`
		FatherStatus            int       `json:"father_status"`
		MotherName              string    `json:"mother_name"`
		PlaceOfBirthMother      string    `json:"place_of_birth_mother"`
		BirthOfDateMother       time.Time `json:"birth_of_date_mother"`
		MotherOccupation        string    `json:"mother_occupation"`
		MotherPhone             string    `json:"mother_phone"`
		MotherStatus            int       `json:"mother_status"`
		Image                   string    `json:"image"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
	}
)

func (res *ListStudentResponse) AddDomain(student []domain.Student) {
	response := ListStudentResponseData{}

	for _, std := range student {

		response.ID = std.ID
		response.IdentityNumber = std.IdentityNumber
		response.SchoolName = std.School.Name
		response.Name = std.Name
		response.Age = std.Age
		response.PlaceOfBirth = std.PlaceOfBirth
		response.BirthOfDate = std.BirthOfDate
		response.ChildRow = std.ChildRow
		response.TotalSibling = std.TotalSibling
		response.Address = std.Address
		response.SosialStatus = std.SosialStatus
		response.PosCode = std.PosCode
		response.ProvinceName = std.Province.Name
		response.RegencyName = std.Regency.Name
		response.DistrictName = std.District.Name
		response.VillageName = std.Village.Name
		response.EducationStatus = std.EducationStatus
		response.RegisteredAt = std.RegisteredAt
		response.FinishedAt = std.FinishedAt
		response.PunishmentCount = std.PunishmentCount
		response.PunishmentStart = std.PunishmentStart
		response.PunishmentEnd = std.PunishmentEnd
		response.JuzKuranDescription = std.JuzKuranDescription
		response.ChapterKuranDescription = std.ChapterKuranDescription
		response.HadistDescription = std.HadistDescription
		response.EducationDescription = std.EducationDescription
		response.ParentStatus = std.ParentStatus
		response.FatherName = std.FatherName
		response.PlaceOfBirthFather = std.PlaceOfBirthFather
		response.BirthOfDateFather = std.BirthOfDateFather
		response.FatherOccupation = std.FatherOccupation
		response.FatherPhone = std.FatherPhone
		response.FatherStatus = std.FatherStatus
		response.MotherName = std.MotherName
		response.PlaceOfBirthMother = std.PlaceOfBirthMother
		response.BirthOfDateMother = std.BirthOfDateMother
		response.MotherOccupation = std.MotherOccupation
		response.MotherPhone = std.MotherPhone
		response.MotherStatus = std.MotherStatus
		response.Image = std.Image
		response.CreatedAt = std.CreatedAt
		response.UpdatedAt = std.UpdatedAt

		res.Data = append(res.Data, response)
	}
}

func SetResponse(domain []ListStudentResponseData, pagination base.PaginationResponseData, message string, success bool) ListStudentResponse {
	return ListStudentResponse{
		BaseResponse: base.BaseResponse{
			Success: success,
			Message: message,
		},
		Pagination: pagination,
		Data:       domain,
	}
}

func ResponseMapper(domain []domain.Student) []ListStudentResponseData {
	response := ListStudentResponse{}

	response.AddDomain(domain)
	return response.Data
}
