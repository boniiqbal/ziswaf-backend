package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type StatementCategory struct {
	base.ModelSoftDelete
	Name       string `gorm:"column:name" json:"name"`
	CategoryID uint64 `gorm:"column:category_id" json:"category_id"`
}
