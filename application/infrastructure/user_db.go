package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type UserRepository interface {
	CreateUser(context.Context, domain.User) (domain.User, error)
	ListUsers(context.Context, domain.UserFilter, []domain.Employee) ([]domain.User, int)
	GetUserById(context.Context, string) (domain.User, error)
	GetUserByUsername(context.Context, string) (domain.User, error)
	UpdateUser(context.Context, domain.User, string) (domain.User, error)
	GetUserByEmployeeId(context.Context, uint64) (domain.User, error)
	DeleteUserByID(context.Context, string) error
}
