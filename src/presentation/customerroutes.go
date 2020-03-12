package presentation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/vagner-nascimento/go-poc-archref/environment"
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/src/tool"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Put("/{id}", putCustomer)
	router.Delete("/{id}", deleteCustomer)
	router.Get("/{id}", getCustomer)
	router.Get("/", getCustomersPaginated)

	return router
}

// TODO: validate params (path, query, body, etc)
// TODO: improve httpErrors
// TODO: realise how to send an safe error into response
// TODO: clean duplicate codes
func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes"))
		return
	}

	if customer, err := app.CreateCustomer(bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, customer)
	}
}

func putCustomer(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	id := params[len(params)-1]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "http request", "bytes"))
		return
	}

	if customer, err := app.UpdateCustomer(id, bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, customer)
	}
}

// TODO: implement DELETE CUSTOMER
func deleteCustomer(w http.ResponseWriter, r *http.Request) {

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	if customer, err := app.FindCustomer(chi.URLParam(r, "id"), &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
		return
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
		pageSize int64
	)

	for key := range query {
		if key == "page" {
			page, err = tool.SafeParseInt(query.Get(key))
			if err != nil {
				render.JSON(w, r, simpleError(err, "cant convert page into int"))
				return
			}
		} else if key == "pageSize" {
			pageSize, err = tool.SafeParseInt(query.Get(key))
			if err != nil {
				render.JSON(w, r, simpleError(err, "cant convert pageSize into int"))
				return
			}
		} else {
			param := query.Get(key)

			var paramValues []interface{}
			if tool.StringIsArray(param) {
				dec := json.NewDecoder(strings.NewReader(param))
				if err = dec.Decode(&paramValues); err != nil {
					render.JSON(w, r, simpleError(err, fmt.Sprintf("cant convert %s into query param", key)))
					return
				}
			} else {
				paramValues = append(paramValues, strings.Replace(param, "\"", "", -1))
			}

			params = append(params, app.SearchParameter{
				Field:  key,
				Values: paramValues,
			})
		}
	}

	if pageSize == 0 {
		pageSize = environment.MaxPaginatedSearch
	}

	customers, total, err := app.FindCustomers(params, page, pageSize, &repository.CustomerRepository{})
	if err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, newPaginatedResponse(customers, page, len(customers), total))
	}
}
