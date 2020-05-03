package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type CategoryRepository interface {
	CreateCategory(context.Context, domain.Category) (domain.Category, error)
	GetListCategory(context.Context) []domain.Category
	GetCategoryByID(context.Context, string) (domain.Category, error)
	UpdateCategory(context.Context, domain.Category, string) (domain.Category, error)
	CreateStatementCategory(context.Context, domain.StatementCategory) (domain.StatementCategory, error)
	GetListStatementCategory(context.Context, string) []domain.StatementCategory
	UpdateStatementCategory(context.Context, domain.StatementCategory, string) (domain.StatementCategory, error)
	DeleteStatementCategoryByID(context.Context, string) error
}
