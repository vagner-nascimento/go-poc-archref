package presenter

import (
	"github.com/go-chi/chi"
)

/*
    TODO: Http presenter
	Requirements list:
		From Node Project
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

func HttpRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/customers", newCustomersRoutes()) // TODO call customers routes and try add mount other routers here
	})

	return router
}
