package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type DonorRepository interface {
	CreateDonor(context.Context, domain.Donor, uint64) (domain.Donor, error)
	GetListDonor(context.Context, domain.DonorFilter) ([]domain.Donor, int)
	GetDonorByID(context.Context, string) (domain.Donor, error)
	UpdateDonor(context.Context, domain.Donor, string) (domain.Donor, error)
	GetDonorByNamePhone(context.Context, domain.Donor) (domain.Donor, error)
	DeleteDonorByID(context.Context, string) error
}
