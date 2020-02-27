package presentation

import (
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Get("/{id}", getCustomer)

	return router
}

func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes")) // TODO: improve httpErrors
	}

	if c, err := app.CreateCustomer(bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err) // TODO: realise how to send an safe error into response
	} else {
		render.JSON(w, r, c) // TODO: realise why it send cardHash (shouldn't send)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c := app.Customer{Id: id, Name: "Test Get"}

	render.JSON(w, r, c)
}