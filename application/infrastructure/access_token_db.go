package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type LoginRepository interface {
	Login(context.Context, domain.AccessToken, uint64) (domain.AccessToken, error)
	GetAccessTokenByUserID(context.Context, uint64) bool
	// UpdateAccessToken(context.Context, domain.AccessToken, uint64) (domain.AccessToken, error)
}
