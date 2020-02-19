package presenter

import (
	"encoding/json"
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
	decoder := json.NewDecoder(r.Body)

	var c app.Customer
	if err := decoder.Decode(&c); err != nil {
		render.JSON(w, r, err)
	}
	if err := app.AddCustomer(&c, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, c) // TODO: realise why it send cardHash (should not send)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	c := app.Customer{Id: id, Name: "Test Get"}

	render.JSON(w, r, c)
}
