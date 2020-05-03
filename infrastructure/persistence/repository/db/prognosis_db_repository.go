package db

import (
	"context"
	"errors"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type prognosisRepository struct {
	DB *gorm.DB
}

func NewPrognosisRepository(DB *gorm.DB) infrastructure.PrognosisRepository {
	return &prognosisRepository{
		DB: DB,
	}
}

func (p *prognosisRepository) CreatePrognosis(ctx context.Context, prognosis []domain.TransactionGoal) ([]domain.TransactionGoal, error) {
	progData := []domain.TransactionGoal{}
	tx := p.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return prognosis, err
	}

	for _, prog := range prognosis {
		p.DB.Where("division_id = ? AND year = ? AND month = ?", prog.DivisionID, prog.Year, prog.Month).Find(&progData)

		if prog.DivisionID != 1 && prog.DivisionID != 2 && prog.DivisionID != 3 {
			tx.Rollback()
			return prognosis, errors.New("Kategori sumber tidak ditemukan")
		}

		if len(progData) == 0 {
			if errProg := tx.Create(&prog).Error; errProg != nil {
				tx.Rollback()
				return prognosis, errProg
			}
		} else {
			prognosisData := domain.TransactionGoal{}

			if err := tx.Model(&prognosisData).Where("id = ?", progData[0].ID).Update(map[string]interface{}{
				"DivisionID": prog.DivisionID,
				"Month":      prog.Month,
				"Year":       prog.Year,
				"Total":      prog.Total}).Error; err != nil {
				tx.Rollback()
				return prognosis, err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return prognosis, err
	}

	return prognosis, nil
}

func (p *prognosisRepository) GetListPrognosis(ctx context.Context, filter domain.TransactionGoalFilter) ([]domain.TransactionGoal, int) {
	var (
		divArrayID []uint64
		count      int
	)

	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	_prognosis := []domain.TransactionGoal{}
	_division := []domain.Division{}

	tx := p.DB.Model(&_prognosis)
	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.DivisionID != "" {
		tx = tx.Where("division_id = ?", filter.DivisionID)
	}
	if filter.StartTotal != "" {
		tx = tx.Where("total >= ?", filter.StartTotal)
	}
	if filter.EndTotal != "" {
		tx = tx.Where("total <= ?", filter.EndTotal)
	}
	if filter.StartDate != "" {
		tx = tx.Where("updated_at >= ?", filter.StartDate)
	}
	if filter.EndDate != "" {
		tx = tx.Where("updated_at <= ?", filter.EndDate)
	}
	if filter.Year != "" {
		tx = tx.Where("year = ?", filter.Year)
	}

	if filter.Search != "" {
		p.DB.Where("name LIKE ?", "%"+filter.Search+"%").Find(&_division)
		for _, v := range _division {
			divArrayID = append(divArrayID, v.ID)
		}
		tx = tx.Where("division_id IN (?)", divArrayID)
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Offset(offset).Limit(limit).Order(filter.Sort).Find(&_prognosis)
	} else {
		tx.Order(filter.Sort).Find(&_prognosis)
	}

	return _prognosis, count
}

func (p *prognosisRepository) GetPrognosisByID(ctx context.Context, id string) (domain.TransactionGoal, error) {
	prognosis := domain.TransactionGoal{}

	if p.DB.First(&prognosis, id).RecordNotFound() {
		return prognosis, errors.New("Tidak dapat menemukan prognosis dengan id " + id)
	}

	return prognosis, nil
}

func (p *prognosisRepository) UpdatePrognosis(ctx context.Context, prognosis domain.TransactionGoal, id string) (domain.TransactionGoal, error) {
	prognosisData := domain.TransactionGoal{}
	err := p.DB.Model(&prognosisData).Where("id = ?", id).Update(&prognosis).Error

	if err != nil {
		return prognosis, err
	}

	return prognosis, nil
}
