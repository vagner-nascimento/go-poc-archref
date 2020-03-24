package presentation

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/vagner-nascimento/go-poc-archref/src/tool"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Put("/{id}", putCustomer)
	//router.Patch("/{id}/email", patchCustomer)
	//router.Delete("/{id}", deleteCustomer)
	router.Get("/{id}", getCustomer)
	router.Get("/", getCustomersPaginated)
	return router
}

func getIdFromPath(path string, skip int) string {
	params := strings.Split(path, "/")
	return params[len(params)-skip]
}

// TODO: validate params (path, query, body, etc)
// TODO: improve httpErrors
// TODO: realise how to send an safe error into response
// TODO: clean duplicate codes
func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	customer, err := model.NewCustomerFromJsonBytes(bytes)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	customerUc := provider.CustomerUseCase()
	if customerUc.Create(&customer); err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, customer)
	}
}

func putCustomer(w http.ResponseWriter, r *http.Request) {
	id := getIdFromPath(r.URL.Path, 1)
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	customer, err := model.NewCustomerFromJsonBytes(bytes)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	customerUs := provider.CustomerUseCase()
	if customer, err := customerUs.Update(id, customer); err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, customer)
	}
}

// TODO: implement patch
func patchCustomer(w http.ResponseWriter, r *http.Request) {

}

// TODO: implement DELETE CUSTOMER
func deleteCustomer(w http.ResponseWriter, r *http.Request) {

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	customerUs := provider.CustomerUseCase()
	if customer, err := customerUs.Find(chi.URLParam(r, "id")); err != nil {
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
		params   []model.SearchParameter
		page     int64
		pageSize int64
	)

	for key := range query {
		if key == "page" {
			page, err = tool.SafeParseInt(query.Get(key))
			if err != nil {
				render.JSON(w, r, err)
				return
			}
		} else if key == "pageSize" {
			pageSize, err = tool.SafeParseInt(query.Get(key))
			if err != nil {
				render.JSON(w, r, err)
				return
			}
		} else {
			param := query.Get(key)
			var paramValues []interface{}
			if tool.IsStringArray(param) {
				dec := json.NewDecoder(strings.NewReader(param))
				if err = dec.Decode(&paramValues); err != nil {
					render.JSON(w, r, err)
					return
				}
			} else {
				paramValues = append(paramValues, strings.Replace(param, "\"", "", -1))
			}

			params = append(params, model.SearchParameter{
				Field:  key,
				Values: paramValues,
			})
		}
	}

	if pageSize == 0 {
		pageSize = config.Get().Data.NoSql.Mongo.MaxPaginatedSearch
	}

	customerUs := provider.CustomerUseCase()
	if customers, total, err := customerUs.List(params, page, pageSize); err != nil {
		render.JSON(w, r, err)
		return
	} else {
		render.JSON(w, r, newPaginatedResponse(customers, page, len(customers), total))
	}
}
