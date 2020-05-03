package entities

import (
	"time"
	base "ziswaf-backend/domain/infrastructure"
)

// School struct models
type School struct {
	base.ModelSoftDelete
	Name        string       `gorm:"column:name" json:"name"`
	Phone       string       `gorm:"column:phone" json:"phone"`
	Email       string       `gorm:"column:email" json:"email"`
	Address     string       `gorm:"column:address" json:"address"`
	PosCode     int          `gorm:"column:pos_code" json:"pos_code"`
	Description string       `gorm:"column:description" json:"description"`
	ProvinceID  uint64       `gorm:"column:province_id" json:"province_id"`
	RegencyID   uint64       `gorm:"column:regency_id" json:"regency_id"`
	HeadMaster  HeadMaster   `json:"head_master"`
	Report      ReportSchool `json:"report_school"`
	Employee    []Employee   `json:"employee"`
	OpenedAt    time.Time    `gorm:"column:opened_at" json:"opened_at"`
	Province    Province     `json:"province"`
	Regency     Regency      `json:"regency"`
	CreatedBy   string       `json:"created_by"`
	UpdatedBy   string       `json:"update_by"`
}

type SchoolFilter struct {
	Search      string `json:"search"`
	Province    string `json:"province"`
	Regency     string `json:"regency"`
	Sort        string `json:"sort"`
	Transaction string `json:"transaction"`
	Page        string `json:"page"`
	Limit       string `json:"limit"`
	Detail      string `json:"detail"`
}

type ReportSchool struct {
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	TotalDonation       int    `json:"total_donation"`
	TotalEmployee       int    `json:"total_employee"`
	TotalTeacher        int    `json:"total_teacher"`
	TotalSosialEmployee int    `json:"total_social_employee"`
	TotalStudent        int    `json:"total_student"`
}

type HeadMaster struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Image string `json:"image"`
}
