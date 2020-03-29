package presentation

import (
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

func StartHttpServer() error {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/customers", newCustomersRoutes())
		r.Mount("/suppliers", newSupplierRoutes())
	})

	walkThroughRoutes := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		logger.Info(fmt.Sprintf("%s %s", method, route))
		return nil
	}

	if err := chi.Walk(router, walkThroughRoutes); err != nil {
		return err
	}

	port := config.Get().Presentation.Web.Port
	logger.Info(fmt.Sprintf("http server listening at port: %d", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
