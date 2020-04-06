package httppresentation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type httpError struct {
	Message *string `json:"message"`
	Type    *string `json:"type"`
	Field   *string `json:"field"`
}

type httpErrors struct {
	Errors []httpError `json:"errors"`
}

func newConversionError(err error) httpError {
	msg := err.Error()
	tp := "validation"

	return httpError{
		Message: &msg,
		Type:    &tp,
		Field:   nil,
	}
}

func newInternalServerError() httpErrors {
	msg := "An unexpected error occurred during request processing"
	tp := "unexpected error"

	return httpErrors{Errors: []httpError{{
		Message: &msg,
		Type:    &tp,
		Field:   nil,
	}}}
}

// TODO: Improve validate body message
// TODO: Realise how to make validate fails when came unexpected fields
func newValidationErrors(errs validator.ValidationErrors) (httpErrs httpErrors) {
	for _, e := range errs {
		msg := fmt.Sprint(e)
		tp := "validation"

		httpErrs.Errors = append(httpErrs.Errors, httpError{
			Message: &msg,
			Type:    &tp,
			Field:   nil,
		})
	}

	return httpErrs
}
