package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type VillageRepository interface {
	GetListVillage(context.Context, string, string) []domain.Village
}