package httppresentation

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"net/http"
)

/*
    TODO: Http presentation - list of tasks bellow
	Requirements list:
		From Node Projects:
		- .use(methodOverride('X-HTTP-Method-Override'))
        - .use(cors())
        - .use(bodyParser.json())
        - .use(compression())
		- express.use(helmet());
        - express.use(helmet.noCache());

		Others:
		- Content type JSON
		- Compress
		- Panic recover to avoid crash server
		- Log api calls
		- Http 2 support (ssl)
*/
//TODO: realise how to validate URL and Query params
func StartServer(errCh chan error) {
	router := chi.NewRouter()
	router.Use(getMiddlewareList()...)
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/customers", newCustomersRoutes())
		r.Mount("/suppliers", newSupplierRoutes())
	})

	walkThroughRoutes := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		logger.Info(fmt.Sprintf("%s %s", method, route))
		return nil
	}

	if err := chi.Walk(router, walkThroughRoutes); err != nil {
		logger.Error("error on verify http routes", err)
		errCh <- errors.New("an error occurred on try to start http server")
		close(errCh)
	} else {
		port := config.Get().Presentation.Web.Port
		logger.Info(fmt.Sprintf("http server connected at port: %d", port))

		errCh <- http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	}
}
