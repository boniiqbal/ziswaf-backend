package db

import (
	"context"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type regencyRepository struct {
	DB *gorm.DB
}

func NewRegencyRepository(DB *gorm.DB) infrastructure.RegencyRepository {
	return &regencyRepository{
		DB: DB,
	}
}

func (p *regencyRepository) GetListRegency(ctx context.Context, search string, provinceID string, school string) []domain.Regency {
	regency := []domain.Regency{}
	schoolData := []domain.School{}
	var regID []uint64

	tx := p.DB.Model(&regency)
	if search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if provinceID != "" {
		tx = tx.Where("province_id = ?", provinceID)
	}
	if school != "" {
		p.DB.Find(&schoolData)
		for _, v := range schoolData {
			regID = append(regID, v.RegencyID)
		}
		tx = tx.Where("id IN (?)", regID)
	}
	if search != "" {
		tx = tx.Where("name LIKE ?", "%"+search+"%")
	}

	tx.Preload("Province").Find(&regency)

	return regency
}
