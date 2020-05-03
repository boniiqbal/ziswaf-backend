package db

import (
	"context"
	"errors"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) infrastructure.CategoryRepository {
	return &categoryRepository{
		DB: DB,
	}
}

func (p *categoryRepository) CreateCategory(ctx context.Context, category domain.Category) (domain.Category, error) {
	err := p.DB.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (p *categoryRepository) GetListCategory(ctx context.Context) []domain.Category {
	categories := []domain.Category{}

	p.DB.Preload("StatementCategories").Find(&categories)

	return categories
}

func (p *categoryRepository) GetCategoryByID(ctx context.Context, id string) (domain.Category, error) {
	category := domain.Category{}

	if p.DB.First(&category, id).RecordNotFound() {
		return category, errors.New("Kategori dengan " + id + " tidak dapat ditemukan")
	}

	return category, nil
}

func (p *categoryRepository) UpdateCategory(ctx context.Context, category domain.Category, id string) (domain.Category, error) {
	categoryData := domain.Category{}
	err := p.DB.Model(&categoryData).Where("id = ?", id).Update(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (p *categoryRepository) CreateStatementCategory(ctx context.Context, stCtgry domain.StatementCategory) (domain.StatementCategory, error) {
	err := p.DB.Create(&stCtgry).Error

	if err != nil {
		return stCtgry, err
	}

	return stCtgry, nil
}

func (p *categoryRepository) GetListStatementCategory(ctx context.Context, filter string) []domain.StatementCategory {
	stCategories := []domain.StatementCategory{}
	trx := []domain.Transaction{}

	tx := p.DB.Model(&stCategories)

	if filter == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter == "1" {
		p.DB.Find(&trx)
		var stArraylID []uint64
		for _, v := range trx {
			stArraylID = append(stArraylID, v.StatementCategoryID)
		}
		tx = tx.Where("id IN (?)", stArraylID)
	}

	tx.Find(&stCategories)

	return stCategories
}

func (p *categoryRepository) UpdateStatementCategory(ctx context.Context, category domain.StatementCategory, id string) (domain.StatementCategory, error) {
	categoryData := domain.StatementCategory{}
	err := p.DB.Model(&categoryData).Where("id = ?", id).Update(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (p *categoryRepository) DeleteStatementCategoryByID(ctx context.Context, id string) error {
	cat := domain.StatementCategory{}
	
	err := p.DB.Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}

	return nil
}
