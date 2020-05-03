package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type DivisionRepository interface {
	CreateDivision(context.Context, domain.Division) (domain.Division, error)
	GetListDivision(context.Context) []domain.Division
	GetDivisionByID(context.Context, string) (domain.Division, error)
	UpdateDivision(context.Context, domain.Division, string) (domain.Division, error)
}