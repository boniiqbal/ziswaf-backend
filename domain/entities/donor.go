package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type Donor struct {
	base.ModelSoftDelete
	Name        string        `gorm:"column:name" json:"name"`
	CompanyName string        `gorm:"column:company_name" json:"company_name"`
	IsCompany   bool          `gorm:"column:is_company" json:"is_company"`
	Position    string        `gorm:"column:position" json:"position"`
	Email       string        `gorm:"column:email" json:"email"`
	Address     string        `gorm:"column:address" json:"address"`
	Phone       string        `gorm:"column:phone" json:"phone"`
	Status      int           `gorm:"column:status" json:"status"`
	Npwp        int64         `gorm:"column:npwp" json:"npwp"`
	PosCode     int           `gorm:"column:pos_code" json:"pos_code"`
	Info        string        `gorm:"column:info" json:"info"`
	ProvinceID  uint64        `gorm:"column:province_id" json:"province_id"`
	RegencyID   uint64        `gorm:"column:regency_id" json:"regency_id"`
	Transaction []Transaction `json:"transaction"`
}

type DonorFilter struct {
	Search        string `json:"search"`
	SchoolID      string `json:"school_id"`
	Regency       string `json:"regency"`
	Status        string `json:"status"`
	DonorCategory string `json:"donor_category"`
	Sort          string `json:"sort"`
	Page          string `json:"page"`
	Limit         string `json:"limit"`
}
