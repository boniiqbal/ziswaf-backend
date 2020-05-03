package entities

import (
	base "ziswaf-backend/domain/infrastructure"
)

type TransactionGoal struct {
	base.Model
	Total        uint64 `gorm:"column:total" json:"total"`
	Month        int    `gorm:"column:month" json:"month"`
	Year         int    `gorm:"column:year" json:"year"`
	DivisionID   uint64 `gorm:"column:division_id" json:"division_id"`
	DivisionName string `gorm:"-"`
}

type TransactionGoalFilter struct {
	Search     string `json:"search"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	StartTotal string `json:"start_total"`
	EndTotal   string `json:"end_total"`
	Sort       string `json:"sort"`
	DivisionID string `json:"division_id"`
	Year       string `json:"year"`
	Page       string `json:"page"`
	Limit      string `json:"limit"`
}
