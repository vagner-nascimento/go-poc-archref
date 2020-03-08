package presentation

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Put("/{id}", putCustomer)
	router.Get("/{id}", getCustomer)
	router.Get("/", getCustomersPaginated)

	return router
}

func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes")) // TODO: improve httpErrors
	}

	if customer, err := app.CreateCustomer(bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err) // TODO: realise how to send an safe error into response
		return
	} else {
		render.JSON(w, r, customer)
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
	id := params[len(params)-1] // TODO: test if params can comes empty
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "http request", "bytes")) // TODO: clean duplicate codes
	}

	if customer, err := app.UpdateCustomer(id, bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, customer)
	}
}

func getCustomersPaginated(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var (
		err      error
		params   []app.SearchParameter
		page     int64
		quantity int64
	)

	for key := range query {
		if key == "page" {
			page, err = strconv.ParseInt(query.Get(key), 0, 64)
			if err != nil {
				render.JSON(w, r, simpleError(err, "cant convert page into int"))
				return
			}
		} else if key == "quantity" {
			quantity, err = strconv.ParseInt(query.Get(key), 0, 64)
			if err != nil {
				render.JSON(w, r, simpleError(err, "cant convert quantity into int"))
				return
			}
		} else {
			params = append(params, app.SearchParameter{
				Field: key,
				Value: query.Get(key),
			})
		}
	}

	if quantity == 0 {
		quantity = 100
	}
	customers, total, err := app.FindCustomers(params, page, quantity, &repository.CustomerRepository{})
	if err != nil { // TODO: handle not found to sent empty array
		render.JSON(w, r, err)
	} else {
		pgQuant := len(customers)
		res := newPaginatedResponse(customers, page, pgQuant, total)
		render.JSON(w, r, res)
	}
}
