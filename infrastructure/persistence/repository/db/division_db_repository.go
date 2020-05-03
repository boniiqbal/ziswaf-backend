package db

import (
	"context"
	"errors"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type divisionRepository struct {
	DB *gorm.DB
}

func NewDivisionRepository(DB *gorm.DB) infrastructure.DivisionRepository {
	return &divisionRepository{
		DB: DB,
	}
}

func (p *divisionRepository) CreateDivision(ctx context.Context, division domain.Division) (domain.Division, error) {
	err := p.DB.Create(&division).Error

	if err != nil {
		return division, err
	}

	return division, nil
}

func (p *divisionRepository) GetListDivision(ctx context.Context) []domain.Division {
	_division := []domain.Division{}

	p.DB.Find(&_division)

	return _division
}

func (p *divisionRepository) GetDivisionByID(ctx context.Context, id string) (domain.Division, error) {
	division := domain.Division{}

	if p.DB.First(&division, id).RecordNotFound() {
		return division, errors.New("Divisi dengan id " + id + " tidak dapat ditemukan")
	}

	return division, nil
}

func (p *divisionRepository) UpdateDivision(ctx context.Context, division domain.Division, id string) (domain.Division, error) {
	divisionData := domain.Division{}
	err := p.DB.Model(&divisionData).Where("id = ?", id).Update(&division).Error

	if err != nil {
		return division, err
	}

	return division, nil
}
