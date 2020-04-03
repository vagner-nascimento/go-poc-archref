package presentation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

//TODO: Improve validate body message
func validateBody(data interface{}) (errs httpErrors) {
	v := validator.New()
	if vErrs := v.Struct(data); vErrs != nil {
		for _, e := range vErrs.(validator.ValidationErrors) {
			msg := fmt.Sprint(e)
			tp := "validation"
			errs.Errors = append(errs.Errors, httpError{
				Message: &msg,
				Type:    &tp,
				Field:   nil,
			})
		}
	}
	return errs
}
