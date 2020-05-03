package infrastructure

import (
	"context"
	domain "ziswaf-backend/domain/entities"
)

type PrognosisRepository interface {
	CreatePrognosis(context.Context, []domain.TransactionGoal) ([]domain.TransactionGoal, error)
	GetListPrognosis(context.Context, domain.TransactionGoalFilter) ([]domain.TransactionGoal, int)
	GetPrognosisByID(context.Context, string) (domain.TransactionGoal, error)
	UpdatePrognosis(context.Context, domain.TransactionGoal, string) (domain.TransactionGoal, error)
}
