package server

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *RequestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
