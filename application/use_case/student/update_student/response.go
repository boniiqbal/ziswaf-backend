package update_student

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateStudentResponse struct {
		base.BaseResponse
		Data UpdateStudentResponseData `json:"data"`
	}

	UpdateStudentResponseData struct {
		ID                      uint64    `json:"id"`
		IdentityNumber          string    `json:"identity_number" form:"identity_number"`
		Name                    string    `json:"name" form:"name"`
		Age                     string    `json:"age" form:"age" validate:"numeric"`
		PlaceOfBirth            string    `json:"place_of_birth" form:"place_of_birth"`
		BirthOfDate             time.Time `json:"birth_of_date" form:"birth_of_date"`
		ChildRow                string    `json:"child_row" form:"child_row"`
		TotalSibling            string    `json:"total_sibling" form:"total_sibling"`
		Address                 string    `json:"address" form:"address"`
		SosialStatus            int       `json:"sosial_status" form:"sosial_status"`
		PosCode                 int       `json:"pos_code" form:"pos_code"`
		ProvinceID              uint64    `json:"province_id" form:"province_id"`
		RegencyID               uint64    `json:"regency_id" form:"regency_id"`
		DistrictID              uint64    `json:"district_id" form:"district_id"`
		VillageID               uint64    `json:"village_id" form:"village_id"`
		EducationStatus         int       `json:"education_status" form:"education_status"`
		RegisteredAt            time.Time `json:"registered_at" form:"registered_at"`
		FinishedAt              time.Time `json:"finished_at" form:"finished_at"`
		PunishmentCount         int       `json:"punishment_count" form:"punishment_count"`
		PunishmentStart         time.Time `json:"punishment_start" form:"punishment_start"`
		PunishmentEnd           time.Time `json:"punishment_end" form:"punishment_end"`
		JuzKuranDescription     string    `json:"juz_kuran_description" form:"juz_kuran_description"`
		ChapterKuranDescription string    `json:"chapter_kuran_description" form:"chapter_kuran_description"`
		HadistDescription       string    `json:"hadist_description" form:"hadist_description"`
		EducationDescription    string    `json:"education_description" form:"education_description"`
		ParentStatus            int       `json:"parent_status" form:"parent_status"`
		FatherName              string    `json:"father_name" form:"father_name"`
		PlaceOfBirthFather      string    `json:"place_of_birth_father" form:"place_of_birth_father"`
		BirthOfDateFather       time.Time `json:"birth_of_date_father" form:"birth_of_date_father"`
		FatherOccupation        string    `json:"father_occupation" form:"father_occupation"`
		FatherPhone             string    `json:"father_phone" form:"father_phone"`
		FatherStatus            int       `json:"father_status" form:"father_status"`
		MotherName              string    `json:"mother_name" form:"mother_name"`
		PlaceOfBirthMother      string    `json:"place_of_birth_mother" form:"place_of_birth_mother"`
		BirthOfDateMother       time.Time `json:"birth_of_date_mother" form:"birth_of_date_mother"`
		MotherOccupation        string    `json:"mother_occupation" form:"mother_occupation"`
		MotherPhone             string    `json:"mother_phone" form:"mother_phone"`
		MotherStatus            int       `json:"mother_status" form:"mother_status"`
		Image                   string    `json:"image"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
	}
)

func SetResponse(domain UpdateStudentResponseData, message string, success bool) UpdateStudentResponse {
	return UpdateStudentResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Student) UpdateStudentResponseData {
	return UpdateStudentResponseData{
		ID:                      domain.ModelSoftDelete.ID,
		IdentityNumber:          domain.IdentityNumber,
		Name:                    domain.Name,
		Age:                     domain.Age,
		PlaceOfBirth:            domain.PlaceOfBirth,
		BirthOfDate:             domain.BirthOfDate,
		ChildRow:                domain.ChildRow,
		TotalSibling:            domain.TotalSibling,
		Address:                 domain.Address,
		SosialStatus:            domain.SosialStatus,
		PosCode:                 domain.PosCode,
		ProvinceID:              domain.ProvinceID,
		RegencyID:               domain.RegencyID,
		DistrictID:              domain.DistrictID,
		VillageID:               domain.VillageID,
		EducationStatus:         domain.EducationStatus,
		RegisteredAt:            domain.RegisteredAt,
		FinishedAt:              domain.FinishedAt,
		PunishmentCount:         domain.PunishmentCount,
		PunishmentStart:         domain.PunishmentStart,
		PunishmentEnd:           domain.PunishmentEnd,
		JuzKuranDescription:     domain.JuzKuranDescription,
		ChapterKuranDescription: domain.ChapterKuranDescription,
		HadistDescription:       domain.HadistDescription,
		EducationDescription:    domain.EducationDescription,
		ParentStatus:            domain.ParentStatus,
		FatherName:              domain.FatherName,
		PlaceOfBirthFather:      domain.PlaceOfBirthFather,
		BirthOfDateFather:       domain.BirthOfDateFather,
		FatherOccupation:        domain.FatherOccupation,
		FatherPhone:             domain.FatherPhone,
		FatherStatus:            domain.FatherStatus,
		MotherName:              domain.MotherName,
		PlaceOfBirthMother:      domain.PlaceOfBirthMother,
		BirthOfDateMother:       domain.BirthOfDateMother,
		MotherOccupation:        domain.MotherOccupation,
		MotherPhone:             domain.MotherPhone,
		MotherStatus:            domain.MotherStatus,
		Image:                   domain.Image,
		CreatedAt:               domain.ModelSoftDelete.CreatedAt,
		UpdatedAt:               domain.ModelSoftDelete.UpdatedAt,
	}
}
