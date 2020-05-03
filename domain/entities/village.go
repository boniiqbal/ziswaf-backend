package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

// Village struct models
type Village struct {
	base.Model
	Name       string `json:"name"`
	DistrictID uint64 `json:"district_id"`
}
