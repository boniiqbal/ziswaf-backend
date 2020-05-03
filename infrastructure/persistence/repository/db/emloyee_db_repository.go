package db

import (
	"context"
	"errors"
	"strconv"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type employeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) infrastructure.EmployeeRepository {
	return &employeeRepository{
		DB: DB,
	}
}

func (p *employeeRepository) CreateEmployee(ctx context.Context, employee domain.Employee, imageFile string) (domain.Employee, error) {
	var (
		province     domain.Province
		regency      domain.Regency
		employeeData []domain.Employee
		school       domain.School
	)

	if p.DB.First(&school, employee.SchoolID).RecordNotFound() {
		return employee, errors.New("Ma'had tidak ditemukan")
	}

	if employee.Status == 1 {
		p.DB.Where("school_id = ? AND status = ?", employee.SchoolID, 1).Find(&employeeData)
		if len(employeeData) != 0 {
			return employee, errors.New("Kepala sekolah sudah terdaftar")
		}
	}

	err := p.DB.Where("id = ?", employee.ProvinceID).First(&province).Error
	if err != nil {
		return employee, errors.New("Provinsi tidak dapat ditemukan")
	}

	if p.DB.Where("province_id = ?", employee.ProvinceID).First(&regency, employee.RegencyID).RecordNotFound() {
		return employee, errors.New("Wilayah tidak dapat ditemukan")
	}

	employee.Image = imageFile

	errCreate := p.DB.Create(&employee).Error
	if errCreate != nil {
		return employee, errCreate
	}

	return employee, nil
}

func (p *employeeRepository) GetListEmployee(ctx context.Context, filter domain.EmployeeFilter) ([]domain.Employee, int) {
	_employee := []domain.Employee{}
	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	var count int

	tx := p.DB.Model(&_employee)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.Regency != "" {
		tx = tx.Where("regency_id = ?", filter.Regency)
	}
	if filter.Province != "" {
		tx = tx.Where("province_id = ?", filter.Province)
	}
	if filter.RegisterStart != "" {
		tx = tx.Where("registered_year >= ?", filter.RegisterStart)
	}
	if filter.RegisterEnd != "" {
		tx = tx.Where("registered_year <= ?", filter.RegisterEnd)
	}
	if filter.SchoolID != "" {
		tx = tx.Where("school_id = ?", filter.SchoolID)
	}
	if filter.Status != "" {
		tx = tx.Where("status = ?", filter.Status)
	}

	if filter.Search != "" {
		tx = tx.Where("name LIKE ? OR phone LIKE ?", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Offset(offset).Limit(limit).Preload("School").Preload("Province").Preload("Regency").Order(filter.Sort).Find(&_employee)
	} else {
		tx.Order(filter.Sort).Preload("School").Preload("Province").Preload("Regency").Find(&_employee)
	}

	return _employee, count
}

func (p *employeeRepository) GetEmployeeByID(ctx context.Context, id string) (domain.Employee, error) {
	employee := domain.Employee{}

	if p.DB.Preload("Province").Preload("Regency").Preload("School").First(&employee, id).RecordNotFound() {
		return employee, errors.New("Tidak dapat menemukan personel dengan id " + id)
	}

	return employee, nil
}

func (p *employeeRepository) UpdateEmployee(ctx context.Context, employee domain.Employee, id string, image string, emploData domain.Employee) (domain.Employee, error) {
	employeeData := domain.Employee{}
	empData := []domain.Employee{}
	province := domain.Province{}
	regency := domain.Regency{}

	employee.Image = image
	provID := strconv.FormatUint(employee.ProvinceID, 10)

	if employee.ProvinceID != 0 {
		if p.DB.First(&province, employee.ProvinceID).RecordNotFound() {
			return employee, errors.New("Provinsi dengan id " + provID + " tidak dapat ditemukan")
		}
	} else {
		employee.ProvinceID = emploData.ProvinceID
	}

	if employee.RegencyID != 0 {
		if p.DB.Where("province_id = ?", employee.ProvinceID).First(&regency, employee.RegencyID).RecordNotFound() {
			return employee, errors.New("Tidak dapat menemukan kota di provinsi dengan id " + provID)
		}
	} else {
		employee.RegencyID = emploData.RegencyID
	}

	if employee.Status == 1 {
		p.DB.Where("school_id = ? AND status = ?", emploData.SchoolID, 1).Find(&empData)
		if len(empData) != 0 {
			return employee, errors.New("Kepala sekolah sudah terdaftar, Mohon untuk status kepala sekolah sebelumnya diubah terlebih dahulu")
		}
	}

	err := p.DB.Model(&employeeData).Where("id = ?", id).Update(map[string]interface{}{
		"Name":           employee.Name,
		"PlaceOfBirth":   employee.PlaceOfBirth,
		"BirthOfDate":    employee.BirthOfDate,
		"Phone":          employee.Phone,
		"Email":          employee.Email,
		"Address":        employee.Address,
		"Status":         employee.Status,
		"RegisteredYear": employee.RegisteredYear,
		"PosCode":        employee.PosCode,
		"ProvinceID":     employee.ProvinceID,
		"RegencyID":      employee.RegencyID,
		"Image":          employee.Image,
	}).Error

	if err != nil {
		return employee, err
	}

	return employee, nil

}

func (p *employeeRepository) DeleteEmployeeByID(ctx context.Context, id string) error {
	employee := domain.Employee{}
	user := domain.User{}

	tx := p.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&employee).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("employee_id = ?", id).Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
