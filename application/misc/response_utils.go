package misc

import (
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

func GetErrorStatusCode(err int) int {
	if err == 200 {
		return http.StatusOK
	}

	// logrus.Error(err)
	switch err {
	case 400:
		return http.StatusBadRequest
	case 401:
		return http.StatusUnauthorized
	case 403:
		return http.StatusForbidden
	case 404:
		return http.StatusNotFound
	case 409:
		return http.StatusConflict
	case 422:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

// NewValidatorError validate error
func NewValidatorError(err error) string {
	var errors string

	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "min":
			errors = v.Field() + " Minimum " + v.Param()
		case "max":
			errors = v.Field() + " Maximum " + v.Param()
		case "email":
			errors = v.Field() + " must be a valid email "
		case "numeric":
			errors = v.Field() + " must be a valid number"
		case "required":
			errors = v.Field() + " cannot be empty "
		}
	}

	return errors
}
