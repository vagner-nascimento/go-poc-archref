package presentation

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Get("/{id}", getCustomer)
	router.Put("/{id}", putCustomer)

	return router
}

func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes")) // TODO: improve httpErrors
	}

	if customer, err := app.CreateCustomer(bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err) // TODO: realise how to send an safe error into response
	} else {
		render.JSON(w, r, customer) // TODO: realise why it send cardHash (shouldn't send)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	if customer, err := app.FindCustomer(chi.URLParam(r, "id"), &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, customer)
	}
}

func putCustomer(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes")) // TODO: clean duplicate codes
	}

	if customer, err := app.UpdateCustomer(id, bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, customer)
	}
}
