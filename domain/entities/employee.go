package entities

import (
	"time"
	base "ziswaf-backend/domain/infrastructure"
)

type Employee struct {
	base.ModelSoftDelete
	SchoolID       uint64    `json:"school_id"`
	Name           string    `json:"name"`
	PlaceOfBirth   string    `json:"place_of_birth"`
	BirthOfDate    time.Time `json:"birth_of_date"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Address        string    `json:"address"`
	Status         int       `json:"status"`
	RegisteredYear time.Time    `json:"registered_year"`
	PosCode        int       `json:"pos_code"`
	ProvinceID     uint64    `json:"province_id"`
	RegencyID      uint64    `json:"regency_id"`
	Image          string    `json:"image"`
	School         School    `json:"school"`
	Province       Province  `json:"province"`
	Regency        Regency   `json:"regency"`
}

type EmployeeFilter struct {
	Search        string `json:"search"`
	RegisterStart string `json:"register_start"`
	RegisterEnd   string `json:"register_end"`
	SchoolID      string `json:"school_id"`
	Status        string `json:"status"`
	Province      string `json:"province"`
	Regency       string `json:"regency"`
	Sort          string `json:"sort"`
	Page          string `json:"page"`
	Limit         string `json:"limit"`
}
