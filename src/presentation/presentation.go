package presentation

import httppresentation "github.com/vagner-nascimento/go-poc-archref/src/presentation/http"

func StartHttpPresentation() error {
	return httppresentation.StartServer()
}
