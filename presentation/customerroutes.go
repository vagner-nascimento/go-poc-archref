package presentation

import (
	"encoding/json"
	"fmt"
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

// TODO: validate params (path, query, body, etc)
// TODO: improve httpErrors
// TODO: realise how to send an safe error into response
// TODO: clean duplicate codes
func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "io.Reader", "bytes"))
	}

	if customer, err := app.CreateCustomer(bytes, &repository.CustomerRepository{}); err != nil {
		render.JSON(w, r, err)
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
	id := params[len(params)-1]
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, castError(err, "http request", "bytes"))
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
	isArray := func(param string) bool {
		return strings.HasPrefix(param, "[") && strings.HasSuffix(param, "]")
	}

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
		} else { // TODO: fix: http Query(): brings only the first parameter. If pass two equal params like this: ?name=Va&name=Je
			jsonParam := query.Get(key)
			var paramValues []interface{}
			if isArray(jsonParam) {
				dec := json.NewDecoder(strings.NewReader(jsonParam))
				if err = dec.Decode(&paramValues); err != nil {
					render.JSON(w, r, simpleError(err, fmt.Sprintf("cant convert %s into query param", key)))
					return
				}
			} else {
				// if strings.HasPrefix(jsonParam, "\"") && strings.HasSuffix(jsonParam, "\"") {
				// 	jsonParam = jsonParam[:len(jsonParam)-len("\"")]
				// 	jsonParam = strings.Replace(jsonParam, "\"", "", 1)
				// }

				paramValues = append(paramValues, strings.Replace(jsonParam, "\"", "", -1))
			}

			params = append(params, app.SearchParameter{
				Field:  key,
				Values: paramValues,
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
		pgQtd := len(customers)
		res := newPaginatedResponse(customers, page, pgQtd, total)
		render.JSON(w, r, res)
	}
}
