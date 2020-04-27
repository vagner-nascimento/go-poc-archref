package presentation

import httppresentation "github.com/vagner-nascimento/go-poc-archref/src/presentation/http"

func StartHttpPresentation(errsCh *chan error) {
	httppresentation.StartServer(errsCh)
}
