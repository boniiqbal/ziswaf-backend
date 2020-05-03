package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type RegencyRepository interface {
	GetListRegency(context.Context, string, string, string) []domain.Regency
}