package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type Goods struct {
	base.Model
	Quantity      int32         `gorm:"column:quantity" json:"quantity"`
	Description   string        `gorm:"column:description" json:"description"`
	Status        int8          `gorm:"column:status" json:"status"`
	CategoryID    int8          `gorm:"column:category_id" json:"category_id"`
	GoodsCategory GoodsCategory `gorm:"foreignkey:CategoryID" json:"goods_categories"`
	Transaction   Transaction   `gorm:"polymorphic:Item;"`
}

type GoodsCategory struct {
	base.Model
	Name string `gorm:"column:name" json:"name"`
}
