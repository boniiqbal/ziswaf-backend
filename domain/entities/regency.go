package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

// Regency struct models
type Regency struct {
	base.Model
	Name         string     `json:"name"`
	ProvinceID   uint64     `json:"province_id"`
	ProvinceName string     `gorm:"-"`
	Province     Province   `json:"province"`
	District     []District `json:"district"`
}
