package db

import (
	"context"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type districtRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(DB *gorm.DB) infrastructure.DistrictRepository {
	return &districtRepository{
		DB: DB,
	}
}

func (p *districtRepository) GetListDistrict(ctx context.Context, search string, regencyID string) []domain.District {
	_district := []domain.District{}

	tx := p.DB.Model(&_district)

	if search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if regencyID != "" {
		tx = tx.Where("regency_id = ?", regencyID)
	}
	if search != "" {
		tx = tx.Where("name LIKE ?", "%"+search+"%")
	}

	tx.Find(&_district)

	return _district
}
