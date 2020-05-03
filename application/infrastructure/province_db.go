package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type ProvinceRepository interface {
	GetListProvince(context.Context, string, string) []domain.Province
	GetProvinceByID(context.Context, string) (domain.Province, error)
}
