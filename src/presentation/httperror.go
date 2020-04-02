package presentation

type httpError struct {
	Message *string `json:"message"`
	Type    *string `json:"type"`
	Field   *string `json:"field"`
}

type httpErrors struct {
	Errors []httpError `json:"errors"`
}
