package presentation

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
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
		r.Mount("/customers", newCustomersRoutes()) // TODO: try mount other routers here
	})

	walkThroughRoutes := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		infra.LogInfo(fmt.Sprintf("%s %s", method, route))
		return nil
	}

	if err := chi.Walk(router, walkThroughRoutes); err != nil {
		return err
	}

	go http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Presentation.Web.Port), router)
	return nil
}
