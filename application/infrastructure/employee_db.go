package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type EmployeeRepository interface {
	CreateEmployee(context.Context, domain.Employee, string) (domain.Employee, error)
	GetListEmployee(context.Context, domain.EmployeeFilter) ([]domain.Employee, int)
	GetEmployeeByID(context.Context, string) (domain.Employee, error)
	UpdateEmployee(context.Context, domain.Employee, string, string, domain.Employee) (domain.Employee, error)
	DeleteEmployeeByID(context.Context, string) error
}
