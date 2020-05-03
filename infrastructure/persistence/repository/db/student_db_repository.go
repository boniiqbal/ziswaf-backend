package db

import (
	"context"
	"errors"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type studentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(DB *gorm.DB) infrastructure.StudentRepository {
	return &studentRepository{
		DB: DB,
	}
}

func (p *studentRepository) CreateStudent(ctx context.Context, student domain.Student, image string) (domain.Student, error) {
	studentData := domain.Student{}
	province := domain.Province{}
	regency := domain.Regency{}
	school := domain.School{}

	if p.DB.First(&school, student.SchoolID).RecordNotFound() {
		return student, errors.New("Ma'had tidak ditemukan")
	}

	provID := strconv.FormatUint(student.ProvinceID, 10)

	if p.DB.First(&province, student.ProvinceID).RecordNotFound() {
		return student, errors.New("Provinsi dengan id " + provID + " tidak dapat ditemukan")
	}

	if p.DB.Where("province_id = ?", student.ProvinceID).First(&regency, student.RegencyID).RecordNotFound() {
		return student, errors.New("Tidak dapat menemukan kota di provinsi dengan id " + provID)
	}

	p.DB.Where("identity_number = ?", student.IdentityNumber).First(&studentData)

	if studentData.IdentityNumber == student.IdentityNumber {
		return student, errors.New("Duplikasi NIS : " + student.IdentityNumber)
	}

	student.Image = image
	err := p.DB.Create(&student).Error

	if err != nil {
		return student, err
	}

	return student, nil
}

func (p *studentRepository) GetListStudent(ctx context.Context, filter domain.StudentFilter) ([]domain.Student, int) {
	_student := []domain.Student{}
	studentFilter := []domain.Student{}
	school := []domain.School{}
	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	var (
		count         int
		schoolArrayID []uint64
	)

	tx := p.DB.Model(&_student)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}

	if filter.Province != "" {
		p.DB.Find(&studentFilter)
		for _, v := range studentFilter {
			schoolArrayID = append(schoolArrayID, v.SchoolID)
		}
		p.DB.Where("id IN (?) AND province_id = ?", schoolArrayID, filter.Province).Find(&school)
		schoolArrayID = []uint64{}
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}

		tx = tx.Where("school_id IN (?)", schoolArrayID)
	}

	if filter.Regency != "" {
		p.DB.Find(&studentFilter)
		for _, v := range studentFilter {
			schoolArrayID = append(schoolArrayID, v.SchoolID)
		}
		p.DB.Where("id IN (?) AND regency_id = ?", schoolArrayID, filter.Regency).Find(&school)
		schoolArrayID = []uint64{}
		for _, v := range school {
			schoolArrayID = append(schoolArrayID, v.ID)
		}

		tx = tx.Where("school_id IN (?)", schoolArrayID)
	}

	if filter.AgeStart != "" {
		tx = tx.Where("age >= ?", filter.AgeStart)
	}
	if filter.AgeEnd != "" {
		tx = tx.Where("age <= ?", filter.AgeEnd)
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
	}
	if filter.RegisteredStart != "" {
		tx = tx.Where("registered_at >= ?", filter.RegisteredStart)
	}
	if filter.RegisteredEnd != "" {
		tx = tx.Where("registered_at <= ?", filter.RegisteredEnd)
	}
	if filter.SosialStatus != "" {
		tx = tx.Where("sosial_status = ?", filter.SosialStatus)
	}
	if filter.EducationStatus != "" {
		tx = tx.Where("education_status = ?", filter.EducationStatus)
	}

	if filter.Search != "" {
		tx = tx.Where("name LIKE ? OR identity_number LIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Offset(offset).Limit(limit).Preload("School").Preload("Province").Preload("Regency").Preload("District").Preload("Village").Order(filter.Sort).Find(&_student)
	} else {
		tx.Preload("School").Preload("Province").Preload("Regency").Preload("District").Preload("Village").Order(filter.Sort).Find(&_student)
	}

	return _student, count
}

func (p *studentRepository) GetStudentByID(ctx context.Context, id string) (domain.Student, error) {
	student := domain.Student{}

	if p.DB.Preload("Province").Preload("Regency").Preload("District").Preload("Village").Preload("School").First(&student, id).RecordNotFound() {
		return student, errors.New("Tidak dapat menemukan siswa dengan id " + id)
	}

	return student, nil
}

func (p *studentRepository) UpdateStudent(ctx context.Context, student domain.Student, id string, image string) (domain.Student, error) {
	studentData := domain.Student{}
	province := domain.Province{}
	regency := domain.Regency{}

	provID := strconv.FormatUint(student.ProvinceID, 10)

	if p.DB.First(&province, student.ProvinceID).RecordNotFound() {
		return student, errors.New("Provinsi dengan id " + provID + " tidak dapat ditemukan")
	}

	if p.DB.Where("province_id = ?", student.ProvinceID).First(&regency, student.RegencyID).RecordNotFound() {
		return student, errors.New("Tidak dapat menemukan kota di provinsi dengan id " + provID)
	}

	student.Image = image

	err := p.DB.Model(&studentData).Where("id = ?", id).Update(&student).Error
	if err != nil {
		return student, err
	}

	return student, nil
}

func (p *studentRepository) DeleteStudentByID(ctx context.Context, id string) error {
	student := domain.Student{}
	err := p.DB.Where("id = ?", id).Delete(&student).Error
	if err != nil {
		return err
	}

	return nil
}
