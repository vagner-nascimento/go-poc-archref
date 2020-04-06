package presentation

import httppresentation "github.com/vagner-nascimento/go-poc-archref/src/presentation/http"

func StartHttpPresentation() <-chan error {
	errCh := make(chan error)
	go httppresentation.StartServer(errCh)

	return errCh
}
