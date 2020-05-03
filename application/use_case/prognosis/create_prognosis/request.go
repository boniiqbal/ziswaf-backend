package create_prognosis

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	CreatePrognosisRequest struct {
		Data []PrognosisDataRequest `json:"data"`
	}

	PrognosisDataRequest struct {
		Total      uint64 `json:"total" validate:"required"`
		Month      int    `json:"month" validate:"required"`
		Year       int    `json:"year" validate:"required"`
		DivisionID uint64 `json:"division_id" validate:"required"`
	}
)

func ValidateRequest(req *CreatePrognosisRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreatePrognosisRequest) []domain.TransactionGoal {
	var (
		transGoalArray []domain.TransactionGoal
		transGoal      domain.TransactionGoal
	)

	for _, v := range req.Data {
		transGoal.Month = v.Month
		transGoal.Year = v.Year
		transGoal.Total = v.Total
		transGoal.DivisionID = v.DivisionID

		transGoalArray = append(transGoalArray, transGoal)
	}
	return transGoalArray
}
