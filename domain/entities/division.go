package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type Division struct {
	base.Model
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Status      int        `gorm:"column:status" json:"status"`
	CreatedBy   string     `gorm:"column:created_by" json:"created_by"`
	UpdatedBy   string     `gorm:"column:updated_by" json:"updated_by"`
	Category    []Category `json:"category"`
}
