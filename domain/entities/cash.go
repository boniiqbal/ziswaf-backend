package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type Cash struct {
	base.Model
	TypeID       int8         `gorm:"column:type_id" json:"type_id"`
	Status       int8         `gorm:"column:status" json:"status"`
	CategoryID   int8         `gorm:"column:category_id" json:"category_id"`
	RefNumber    string       `gorm:"column:ref_number" json:"ref_number"`
	Description  string       `gorm:"column:description" json:"description"`
	CashCategory CashCategory `gorm:"foreignkey:CategoryID" json:"cash_categories"`
	Transaction  Transaction  `gorm:"polymorphic:Item;"`
}

type CashCategory struct {
	base.Model

	Name string `gorm:"column:name" json:"name"`
}

func (Cash) TableName() string {
	return "cashs"
}
