package db

import (
	"context"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type villageRepository struct {
	DB *gorm.DB
}

func NewVillageRepository(DB *gorm.DB) infrastructure.VillageRepository {
	return &villageRepository{
		DB: DB,
	}
}

func (p *villageRepository) GetListVillage(ctx context.Context, search string, districtID string) []domain.Village {
	_village := []domain.Village{}

	tx := p.DB.Model(&_village)

	if search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if districtID != "" {
		tx = tx.Where("district_id = ?", districtID)
	}
	if search != "" {
		tx = tx.Where("name LIKE ?", "%"+search+"%")
	}

	tx.Find(&_village)

	return _village
}
