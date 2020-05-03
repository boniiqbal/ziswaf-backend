package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

// Province struct models
type Province struct {
	base.Model
	Name    string    `json:"name"`
	Regency []Regency `json:"regency"`
}
