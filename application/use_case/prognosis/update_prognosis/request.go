package update_prognosis

import (
	domain "ziswaf-backend/domain/entities"

	validator "gopkg.in/go-playground/validator.v9"
)

type (
	UpdatePronosisRequest struct {
		Data struct {
			Total      uint64 `json:"total"`
			Month      int    `json:"month"`
			Year       int    `json:"year"`
			DivisionID uint64 `json:"division_id"`
		}
	}
)

func ValidateRequest(req *UpdatePronosisRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req.Data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req UpdatePronosisRequest) domain.TransactionGoal {
	return domain.TransactionGoal{
		Total:      req.Data.Total,
		Month:      req.Data.Month,
		Year:       req.Data.Year,
		DivisionID: req.Data.DivisionID,
	}
}
