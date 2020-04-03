package presentation

type httpError struct {
	Message *string `json:"message"`
	Type    *string `json:"type"`
	Field   *string `json:"field"`
}

type httpErrors struct {
	Errors []httpError `json:"errors"`
}

func getConversionError(err error) httpError {
	msg := err.Error()
	tp := "validation"
	return httpError{
		Message: &msg,
		Type:    &tp,
		Field:   nil,
	}
}
