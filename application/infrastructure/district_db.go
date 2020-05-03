package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type DistrictRepository interface {
	GetListDistrict(context.Context, string, string) []domain.District
}
