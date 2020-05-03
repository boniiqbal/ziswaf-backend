package db

import (
	"context"
	"errors"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	"ziswaf-backend/application/misc"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type schoolRepository struct {
	DB *gorm.DB
}

func NewSchoolRepository(DB *gorm.DB) infrastructure.SchoolRepository {
	return &schoolRepository{
		DB: DB,
	}
}

func (p *schoolRepository) CreateSchool(ctx context.Context, school domain.School) (domain.School, error) {
	var (
		province domain.Province
		regency  domain.Regency
	)

	provID := strconv.FormatUint(school.ProvinceID, 10)

	err := p.DB.Where("id = ?", school.ProvinceID).First(&province).Error
	if err != nil {
		return school, errors.New("Tidak dapat menemukan provinsi dengan id " + provID)
	}

	if p.DB.Where("province_id = ?", school.ProvinceID).First(&regency, school.RegencyID).RecordNotFound() {
		regID := strconv.FormatUint(school.RegencyID, 10)
		return school, errors.New("Tidak dapat menemukan kota dengan id " + regID + " di provinsi dengan id " + provID)
	}

	errCreate := p.DB.Create(&school).Error
	if errCreate != nil {
		return school, errCreate
	}

	return school, nil
}

func (p *schoolRepository) GetListSchool(ctx context.Context, filter domain.SchoolFilter) ([]domain.School, int) {
	_school := []domain.School{}
	school := []domain.School{}
	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	report := domain.ReportSchool{}
	headMaster := domain.HeadMaster{}
	employee := []domain.Employee{}
	student := []domain.Student{}
	trx := []domain.Transaction{}

	var (
		count               int
		countEmployee       int
		countStudent        int
		totalTeacher        int
		totalSosialEmployee int
	)

	tx := p.DB.Model(&_school)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}

	if filter.Regency != "" {
		regArray := misc.QueryConvert(filter.Regency)
		tx = tx.Where("regency_id IN (?)", regArray)
	}

	if filter.Province != "" {
		provArray := misc.QueryConvert(filter.Province)
		tx = tx.Where("province_id IN (?)", provArray)
	}

	if filter.Search != "" {
		tx = tx.Where("name LIKE ? OR phone LIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	if filter.Transaction == "1" {
		p.DB.Find(&trx)
		var schooArraylID []uint64
		for _, v := range trx {
			schooArraylID = append(schooArraylID, v.SchoolID)
		}
		tx = tx.Where("id IN (?)", schooArraylID)
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Offset(offset).Limit(limit).Preload("Province").Preload("Regency").Order(filter.Sort).Find(&_school)
	} else {
		tx.Order(filter.Sort).Preload("Province").Preload("Regency").Find(&_school)
	}

	if filter.Detail == "" {
		school = _school
	}

	if filter.Detail != "" {
		for _, v := range _school {
			p.DB.Where("school_id = ?", v.ID).Find(&employee).Count(&countEmployee)
			p.DB.Where("school_id = ? AND status IN (?)", v.ID, []int{2, 3}).Find(&employee).Count(&totalTeacher)
			p.DB.Where("school_id = ? AND status = ?", v.ID, 6).Find(&employee).Count(&totalSosialEmployee)
			p.DB.Where("school_id = ? AND status = ?", v.ID, 1).Find(&employee)
			report.TotalEmployee = countEmployee
			report.TotalTeacher = totalTeacher
			report.TotalSosialEmployee = totalSosialEmployee

			if len(employee) == 0 {
				headMaster.ID = 0
				headMaster.Name = ""
				headMaster.Email = ""
				headMaster.Phone = ""
				headMaster.Image = ""
			} else {
				headMaster.ID = employee[0].ID
				headMaster.Name = employee[0].Name
				headMaster.Email = employee[0].Email
				headMaster.Phone = employee[0].Phone
				headMaster.Image = employee[0].Image
			}

			p.DB.Where("school_id = ?", v.ID).Find(&student).Count(&countStudent)
			report.TotalStudent = countStudent

			v.HeadMaster = headMaster
			v.Report = report

			school = append(school, v)
		}
	}

	return school, count
}

func (p *schoolRepository) GetSchoolByID(ctx context.Context, id string) (domain.School, error) {
	_school := domain.School{}
	report := domain.ReportSchool{}
	student := []domain.Student{}
	employee := []domain.Employee{}
	headMaster := domain.HeadMaster{}
	var (
		countEmployee       int
		countStudent        int
		totalTeacher        int
		totalSosialEmployee int
	)

	if p.DB.Preload("Province").Preload("Regency").First(&_school, id).RecordNotFound() {
		return _school, errors.New("Tidak dapat menemukan sekolah dengan id " + id)
	}

	p.DB.Where("school_id = ?", _school.ID).Find(&employee).Count(&countEmployee)
	p.DB.Where("school_id = ? AND status IN (?)", _school.ID, []int{2, 3}).Find(&employee).Count(&totalTeacher)
	p.DB.Where("school_id = ? AND status = ?", _school.ID, 6).Find(&employee).Count(&totalSosialEmployee)
	p.DB.Where("school_id = ? AND status = ?", _school.ID, 1).Find(&employee)

	if len(employee) == 0 {
		headMaster.ID = 0
		headMaster.Name = ""
		headMaster.Email = ""
		headMaster.Phone = ""
		headMaster.Image = ""
	} else {
		headMaster.ID = employee[0].ID
		headMaster.Name = employee[0].Name
		headMaster.Email = employee[0].Email
		headMaster.Phone = employee[0].Phone
		headMaster.Image = employee[0].Image
	}

	report.TotalEmployee = countEmployee
	report.TotalTeacher = totalTeacher
	report.TotalSosialEmployee = totalSosialEmployee

	p.DB.Where("school_id = ?", _school.ID).Find(&student).Count(&countStudent)
	report.TotalStudent = countStudent
	_school.Report = report
	_school.HeadMaster = headMaster

	return _school, nil
}

func (p *schoolRepository) UpdateSchool(ctx context.Context, school domain.School, id string) (domain.School, error) {
	var (
		schoolData domain.School
		province   domain.Province
		regency    domain.Regency
	)

	provID := strconv.FormatUint(school.ProvinceID, 10)

	if p.DB.Where("id = ?", school.ProvinceID).First(&province).RecordNotFound() {
		return school, errors.New("Tidak dapat menemukan provinsi dengan id " + provID)
	}

	if p.DB.Where("province_id = ?", school.ProvinceID).First(&regency, school.RegencyID).RecordNotFound() {
		regID := strconv.FormatUint(school.RegencyID, 10)
		return school, errors.New("Tidak dapat menemukan kota dengan id " + regID + " di provinsi dengan id " + provID)
	}

	err := p.DB.Model(&schoolData).Where("id = ?", id).Update(&school).Error

	if err != nil {
		return school, err
	}

	return school, nil
}

func (p *schoolRepository) GetSchoolByAccountID(ctx context.Context, accountID uint64) (domain.School, error) {
	school := domain.School{}
	if p.DB.Where("user_id = ?", accountID).First(&school).RecordNotFound() {
		return school, errors.New("Sekolah tidak ditemukan pada akun ini")
	}

	return school, nil
}

func (p *schoolRepository) DeleteSchoolByID(ctx context.Context, id string) error {
	school := domain.School{}
	student := []domain.Student{}
	employee := []domain.Employee{}
	trx := []domain.Transaction{}
	user := []domain.User{}

	tx := p.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&school).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("school_id = ?", id).Delete(&student).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("school_id = ?", id).Delete(&employee).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("school_id = ?", id).Delete(&trx).Error; err != nil {
		tx.Rollback()
		return err
	}

	p.DB.Where("school_id = ?", id).Find(&employee)
	if len(employee) != 0 {
		var empArrayID []uint64
		for _, v := range employee {
			empArrayID = append(empArrayID, v.ID)
		}
		if err := tx.Where("employee_id IN (?)", empArrayID).Delete(&user).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (p *schoolRepository) GetRecordSchool(ctx context.Context, id string) (domain.ReportSchool, error) {
	transaction := []domain.Transaction{}
	employee := []domain.Employee{}
	student := []domain.Student{}
	school := domain.School{}

	var (
		transactionRecord int
		employeeRecord    int
		studentRecord     int
	)

	p.DB.Where("id = ?", id).First(&school)
	p.DB.Where("school_id = ?", id).Find(&transaction).Count(&transactionRecord)
	p.DB.Where("school_id = ?", id).Find(&employee).Count(&employeeRecord)
	p.DB.Where("school_id = ?", id).Find(&student).Count(&studentRecord)

	return domain.ReportSchool{
		ID:            school.ModelSoftDelete.ID,
		Name:          school.Name,
		TotalDonation: transactionRecord,
		TotalEmployee: employeeRecord,
		TotalStudent:  studentRecord,
	}, nil
}
