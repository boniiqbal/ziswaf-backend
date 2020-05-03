package db

import (
	"context"
	"errors"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type donorRepository struct {
	DB *gorm.DB
}

func NewDonorRepository(DB *gorm.DB) infrastructure.DonorRepository {
	return &donorRepository{
		DB: DB,
	}
}

func (p *donorRepository) CreateDonor(ctx context.Context, donor domain.Donor, donorID uint64) (domain.Donor, error) {
	province := domain.Province{}
	regency := domain.Regency{}

	provID := strconv.FormatUint(donor.ProvinceID, 10)

	if p.DB.First(&province, donor.ProvinceID).RecordNotFound() {
		return donor, errors.New("Provinsi dengan id " + provID + " tidak dapat ditemukan")
	}

	if p.DB.Where("province_id = ?", donor.ProvinceID).First(&regency, donor.RegencyID).RecordNotFound() {
		return donor, errors.New("Tidak dapat menemukan kota di provinsi dengan id " + provID)
	}

	tx := p.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return donor, err
	}

	if donorID == 0 {
		if err := tx.Create(&donor).Error; err != nil {
			tx.Rollback()
			return donor, err
		}
	} else if donorID != 0 {
		var donorData domain.Donor
		if err := tx.Model(&donorData).Where("id = ?", donorID).Update(&donor).Error; err != nil {
			tx.Rollback()
			return donor, err
		}
		donor.ID = donorID
	}

	if err := tx.Commit().Error; err != nil {
		return donor, err
	}

	return donor, nil
}

func (p *donorRepository) GetListDonor(ctx context.Context, filter domain.DonorFilter) ([]domain.Donor, int) {
	_donor := []domain.Donor{}
	trx := []domain.Transaction{}
	employee := []domain.Employee{}

	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	var (
		count           int
		employeeArrayID []uint64
		donorArrayID    []uint64
	)

	tx := p.DB.Model(&_donor)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.Regency != "" {
		tx = tx.Where("regency_id = ?", filter.Regency)
	}
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
	}
	if filter.DonorCategory != "" {
		tx = tx.Where("is_company = ?", filter.DonorCategory)
	}
	if filter.SchoolID != "" {
		p.DB.Where("school_id = ?", filter.SchoolID).Find(&employee)
		for _, v := range employee {
			employeeArrayID = append(employeeArrayID, v.ID)
		}
		p.DB.Where("employee_id IN (?)", employeeArrayID).Find(&trx)
		for _, v := range trx {
			donorArrayID = append(donorArrayID, v.DonorID)
		}
		tx = tx.Where("id IN (?)", donorArrayID)
	}

	if filter.Search != "" {
		tx = tx.Where("name LIKE ? OR company_name LIKE ? OR phone LIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Offset(offset).Limit(limit).Order(filter.Sort).Find(&_donor)
	} else {
		tx.Order(filter.Sort).Find(&_donor)
	}

	return _donor, count
}

func (p *donorRepository) GetDonorByID(ctx context.Context, id string) (domain.Donor, error) {
	donor := domain.Donor{}

	if p.DB.First(&donor, id).RecordNotFound() {
		return donor, errors.New("Donatur dengan id " + id + " tidak dapat ditemukan")
	}

	return donor, nil
}

func (p *donorRepository) UpdateDonor(ctx context.Context, donor domain.Donor, id string) (domain.Donor, error) {
	donorData := domain.Donor{}
	province := domain.Province{}
	regency := domain.Regency{}

	if donor.ProvinceID != 0 && donor.RegencyID != 0 {
		provID := strconv.FormatUint(donor.ProvinceID, 10)

		if p.DB.First(&province, donor.ProvinceID).RecordNotFound() {
			return donor, errors.New("Provinsi dengan id " + provID + " tidak dapat ditemukan")
		}

		if p.DB.Where("province_id = ?", donor.ProvinceID).First(&regency, donor.RegencyID).RecordNotFound() {
			return donor, errors.New("Tidak dapat menemukan kota di provinsi dengan id " + provID)
		}
	}

	err := p.DB.Model(&donorData).Where("id = ?", id).Update(&donor).Error

	if err != nil {
		return donor, err
	}

	return donor, nil
}

func (p *donorRepository) GetDonorByNamePhone(ctx context.Context, donor domain.Donor) (domain.Donor, error) {
	_donor := domain.Donor{}

	if p.DB.Where("phone = ? AND name = ?", donor.Phone, donor.Name).First(&_donor).RecordNotFound() {
		return _donor, errors.New("No.Telepon sudah tersedia")
	}
	return _donor, nil
}

func (p *donorRepository) DeleteDonorByID(ctx context.Context, id string) error {
	donor := domain.Donor{}
	err := p.DB.Where("id = ?", id).Delete(&donor).Error
	if err != nil {
		return err
	}

	return nil
}
