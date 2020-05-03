package update_student

import (
	"time"
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdateStudentRequest struct {
		Data struct {
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
		}
	}
)

func ValidateRequest(req *UpdateStudentRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdateStudentRequest) domain.Student {
	return domain.Student{
		IdentityNumber:          req.Data.IdentityNumber,
		Name:                    req.Data.Name,
		Age:                     req.Data.Age,
		PlaceOfBirth:            req.Data.PlaceOfBirth,
		BirthOfDate:             req.Data.BirthOfDate,
		ChildRow:                req.Data.ChildRow,
		TotalSibling:            req.Data.TotalSibling,
		Address:                 req.Data.Address,
		SosialStatus:            req.Data.SosialStatus,
		PosCode:                 req.Data.PosCode,
		ProvinceID:              req.Data.ProvinceID,
		RegencyID:               req.Data.RegencyID,
		DistrictID:              req.Data.DistrictID,
		VillageID:               req.Data.VillageID,
		EducationStatus:         req.Data.EducationStatus,
		RegisteredAt:            req.Data.RegisteredAt,
		FinishedAt:              req.Data.FinishedAt,
		PunishmentCount:         req.Data.PunishmentCount,
		PunishmentStart:         req.Data.PunishmentStart,
		PunishmentEnd:           req.Data.PunishmentEnd,
		JuzKuranDescription:     req.Data.JuzKuranDescription,
		ChapterKuranDescription: req.Data.ChapterKuranDescription,
		HadistDescription:       req.Data.HadistDescription,
		EducationDescription:    req.Data.EducationDescription,
		ParentStatus:            req.Data.ParentStatus,
		FatherName:              req.Data.FatherName,
		PlaceOfBirthFather:      req.Data.PlaceOfBirthFather,
		BirthOfDateFather:       req.Data.BirthOfDateFather,
		FatherOccupation:        req.Data.FatherOccupation,
		FatherPhone:             req.Data.FatherPhone,
		FatherStatus:            req.Data.FatherStatus,
		MotherName:              req.Data.MotherName,
		PlaceOfBirthMother:      req.Data.PlaceOfBirthMother,
		BirthOfDateMother:       req.Data.BirthOfDateMother,
		MotherOccupation:        req.Data.MotherOccupation,
		MotherPhone:             req.Data.MotherPhone,
		MotherStatus:            req.Data.MotherStatus,
	}
}
