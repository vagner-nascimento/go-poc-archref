package presentation

import httppresentation "github.com/vagner-nascimento/go-poc-archref/src/presentation/http"

func StartHttpPresentation(errs chan error) {
	httppresentation.StartServer(errs)
}
