package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type SchoolRepository interface {
	CreateSchool(context.Context, domain.School) (domain.School, error)
	GetListSchool(context.Context, domain.SchoolFilter) ([]domain.School, int)
	GetSchoolByID(context.Context, string) (domain.School, error)
	UpdateSchool(context.Context, domain.School, string) (domain.School, error)
	GetSchoolByAccountID(context.Context, uint64) (domain.School, error)
	DeleteSchoolByID(context.Context, string) error
	GetRecordSchool(context.Context, string) (domain.ReportSchool, error)
}
