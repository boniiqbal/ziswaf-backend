package entities

import (
	"time"
	base "ziswaf-backend/domain/infrastructure"
)

// Student struct models
type Student struct {
	base.ModelSoftDelete
	IdentityNumber          string    `json:"identity_number"`
	SchoolID                uint64    `json:"school_id"`
	Name                    string    `json:"name"`
	Age                     string    `json:"age"`
	PlaceOfBirth            string    `json:"place_of_birth"`
	BirthOfDate             time.Time `json:"birth_of_date"`
	ChildRow                string    `json:"child_row"`
	TotalSibling            string    `json:"total_sibling"`
	Address                 string    `json:"address"`
	SosialStatus            int       `json:"sosial_status"`
	PosCode                 int       `json:"pos_code"`
	ProvinceID              uint64    `json:"province_id"`
	RegencyID               uint64    `json:"city_id"`
	DistrictID              uint64    `json:"district_id"`
	VillageID               uint64    `json:"village_id"`
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
	School                  School    `json:"school"`
	Province                Province  `json:"province"`
	Regency                 Regency   `json:"regency"`
	District                District  `json:"district"`
	Village                 Village   `json:"village"`
}

type StudentFilter struct {
	SchoolID        string `json:"school_id"`
	SosialStatus    string `json:"sosial_status"`
	EducationStatus string `json:"education_status"`
	AgeStart        string `json:"age_start"`
	AgeEnd          string `json:"age_end"`
	RegisteredStart string `json:"registered_start"`
	RegisteredEnd   string `json:"registered_end"`
	Search          string `json:"search"`
	Province        string `json:"province"`
	Regency         string `json:"regency"`
	Sort            string `json:"sort"`
	Page            string `json:"page"`
	Limit           string `json:"limit"`
}
