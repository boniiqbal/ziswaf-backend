package db

import (
	"context"
	"errors"
	"strconv"
	"time"
	"ziswaf-backend/application/infrastructure"
	domain "ziswaf-backend/domain/entities"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) infrastructure.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (p *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	employee := domain.Employee{}
	userData := domain.User{}
	user.LastLogin = time.Now()

	// Validate employee id
	err := p.DB.First(&employee, user.EmployeeID).Error
	if err != nil {
		return user, errors.New("Petugas tidak dapat ditemukan")
	}

	// Validate user existing
	errUser := p.DB.Where("employee_id = ?", employee.ID).First(&userData).Error
	if errUser == nil {
		return user, errors.New("Petugas ini sudah didaftarkan")
	}

	user.Name = employee.Name

	errCreate := p.DB.Create(&user).Error

	if errCreate != nil {
		return user, errCreate
	}

	return user, nil
}

func (p *userRepository) ListUsers(ctx context.Context, filter domain.UserFilter, filterEmployee []domain.Employee) ([]domain.User, int) {
	_users := []domain.User{}
	page, _ := strconv.Atoi(filter.Page)
	limit, _ := strconv.Atoi(filter.Limit)
	offset := (page - 1) * limit
	var (
		count           int
		employeeArrayID []uint64
	)

	tx := p.DB.Model(&_users)

	if filter.Search == "" {
		tx = tx.Where("id > ?", 0)
	}
	if filter.CreatedStart != "" {
		tx = tx.Where("created_at >= ?", filter.CreatedStart)
	}
	if filter.CreatedEnd != "" {
		tx = tx.Where("created_at <= ?", filter.CreatedEnd)
	}
	if filter.LoginStart != "" {
		tx = tx.Where("last_login >= ?", filter.LoginStart)
	}
	if filter.LoginEnd != "" {
		tx = tx.Where("last_login <= ?", filter.LoginEnd)
	}
	if filter.Role != "" {
		tx = tx.Where("role = ?", filter.Role)
	}
	if filter.SchoolID != "" {
		for _, v := range filterEmployee {
			employeeArrayID = append(employeeArrayID, v.ID)
		}
		tx = tx.Where("employee_id IN (?)", employeeArrayID)
	}

	if filter.Search != "" {
		for _, v := range filterEmployee {
			employeeArrayID = append(employeeArrayID, v.ID)
		}
		tx = tx.Where("employee_id IN (?)", employeeArrayID)
	}

	tx.Count(&count)

	if filter.Page != "" || filter.Limit != "10" {
		tx.Preload("Employee.School").Offset(offset).Limit(limit).Order(filter.Sort).Find(&_users)
	} else {
		tx.Preload("Employee.School").Order(filter.Sort).Find(&_users)
	}

	return _users, count
}

func (p *userRepository) GetUserById(ctx context.Context, id string) (domain.User, error) {
	user := domain.User{}

	if p.DB.Preload("Employee.School").Preload("AccessToken").First(&user, id).RecordNotFound() {
		return user, errors.New("User dengan ID " + id + " tidak dapat ditemukan")
	}

	return user, nil
}

func (p *userRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	userData := domain.User{}

	if p.DB.Where("username = ?", username).First(&userData).RecordNotFound() {
		return userData, errors.New("Username salah")
	}
	return userData, nil
}

func (p *userRepository) UpdateUser(ctx context.Context, user domain.User, id string) (domain.User, error) {
	userData := domain.User{}
	err := p.DB.Model(&userData).Where("id = ?", id).Update(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (p *userRepository) GetUserByEmployeeId(ctx context.Context, employeeID uint64) (domain.User, error) {
	user := domain.User{}

	if p.DB.Where("employee_id = ?", employeeID).First(&user).RecordNotFound() {
		return user, errors.New("User tidak ditemukan")
	}

	return user, nil
}

func (p *userRepository) DeleteUserByID(ctx context.Context, id string) error {
	user := domain.User{}

	if id == "1" {
		return errors.New("Superuser tidak bisa dihapus")
	}
	
	err := p.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
