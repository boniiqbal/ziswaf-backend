package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type StudentRepository interface {
	CreateStudent(context.Context, domain.Student, string) (domain.Student, error)
	GetListStudent(context.Context, domain.StudentFilter) ([]domain.Student, int)
	GetStudentByID(context.Context, string) (domain.Student, error)
	UpdateStudent(context.Context, domain.Student, string, string) (domain.Student, error)
	DeleteStudentByID(context.Context, string) error
}
