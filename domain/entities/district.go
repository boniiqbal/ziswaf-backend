package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

// District struct models
type District struct {
	base.Model
	Name      string    `json:"name"`
	RegencyID uint64    `json:"regency_id"`
	Village   []Village `json:"village"`
}
