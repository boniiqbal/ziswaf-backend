package create_student

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateEmployeeResponse struct {
		base.BaseResponse
		Data CreateEmployeeResponseData `json:"data"`
	}

	CreateEmployeeResponseData struct {
		ID                      uint64    `json:"id"`
		IdentityNumber          string    `json:"identity_number"`
		SchoolID                uint64    `json:"school_id"`
		Name                    string    `json:"name"`
		Age                     string    `json:"age" validate:"numeric"`
		PlaceOfBirth            string    `json:"place_of_birth"`
		BirthOfDate             time.Time `json:"birth_of_date" `
		ChildRow                string    `json:"child_row"`
		TotalSibling            string    `json:"total_sibling"`
		Address                 string    `json:"address"`
		SosialStatus            int       `json:"sosial_status"`
		PosCode                 int       `json:"pos_code"`
		ProvinceID              uint64    `json:"province_id"`
		RegencyID               uint64    `json:"regency_id"`
		DistrictID              uint64    `json:"district_id"`
		VillageID               uint64    `json:"village_id"`
		EducationStatus         int       `json:"education_status"`
		RegisteredAt            time.Time `json:"registered_at" `
		FinishedAt              time.Time `json:"finished_at"`
		PunishmentCount         int       `json:"punishment_count"`
		PunishmentStart         time.Time `json:"punishment_start" `
		PunishmentEnd           time.Time `json:"punishment_end" `
		JuzKuranDescription     string    `json:"juz_kuran_description"`
		ChapterKuranDescription string    `json:"chapter_kuran_description"`
		HadistDescription       string    `json:"hadist_description"`
		EducationDescription    string    `json:"education_description"`
		ParentStatus            int       `json:"parent_status"`
		FatherName              string    `json:"father_name"`
		PlaceOfBirthFather      string    `json:"place_of_birth_father"`
		BirthOfDateFather       time.Time `json:"birth_of_date_fatheirth_of_date_father"`
		FatherOccupation        string    `json:"father_occupation"`
		FatherPhone             string    `json:"father_phone"`
		FatherStatus            int       `json:"father_status"`
		MotherName              string    `json:"mother_name"`
		PlaceOfBirthMother      string    `json:"place_of_birth_mother"`
		BirthOfDateMother       time.Time `json:"birth_of_date_motheirth_of_date_mother"`
		MotherOccupation        string    `json:"mother_occupation"`
		MotherPhone             string    `json:"mother_phone"`
		MotherStatus            int       `json:"mother_status"`
		Image                   string    `json:"image"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
	}
)

func SetResponse(domain CreateEmployeeResponseData, message string, success bool) CreateEmployeeResponse {
	return CreateEmployeeResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: domain,
	}
}

func ResponseMapper(domain domain.Student) CreateEmployeeResponseData {
	return CreateEmployeeResponseData{
		ID:                      domain.ModelSoftDelete.ID,
		IdentityNumber:          domain.IdentityNumber,
		SchoolID:                domain.SchoolID,
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
