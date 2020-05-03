package db

import (
	"context"
	"errors"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type provinceRepository struct {
	DB *gorm.DB
}

func NewProvinceRepository(DB *gorm.DB) infrastructure.ProvinceRepository {
	return &provinceRepository{
		DB: DB,
	}
}

func (p *provinceRepository) GetListProvince(ctx context.Context, search string, school string) []domain.Province {
	_province := []domain.Province{}
	schoolData := []domain.School{}
	var provinceID []uint64

	tx := p.DB.Model(&_province)

	if search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if school != "" {
		p.DB.Find(&schoolData)
		for _, v := range schoolData {
			provinceID = append(provinceID, v.ProvinceID)
		}
		tx = tx.Where("id IN (?)", provinceID)
	}

	if search != "" {
		tx = tx.Where("name LIKE ?", "%"+search+"%")
	}

	tx.Find(&_province)

	return _province
}

func (p *provinceRepository) GetProvinceByID(ctx context.Context, id string) (domain.Province, error) {
	province := domain.Province{}

	if p.DB.Preload("Regency").First(&province, id).RecordNotFound() {
		return province, errors.New("Provinsi dengan id " + id + " tidak dapat ditemukan")
	}

	return province, nil
}
