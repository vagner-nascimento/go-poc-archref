package presentation

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"io/ioutil"
	"net/http"
)

func newCustomersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postCustomer)
	router.Put("/{id}", putCustomer)
	router.Patch("/{id}/address", patchCustomerAddress)
	//router.Delete("/{id}", deleteCustomer)
	router.Get("/{id}", getCustomer)
	router.Get("/", getCustomersPaginated)
	return router
}

// TODO: validate params (path, query, body, etc)
// TODO: improve http error handler
// TODO: realise how to send an safe error into response
// TODO: clean duplicate codes
func postCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	customer, err := model.NewCustomerFromJsonBytes(bytes) //If int(and i guess that other number too) starts with zero returns error
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		if err = customerUc.Create(&customer); err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, customer)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func putCustomer(w http.ResponseWriter, r *http.Request) {
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
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := customerUc.Update(id, customer); err != nil {
			render.JSON(w, r, err)
			return
		} else {
			render.JSON(w, r, customer)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func patchCustomerAddress(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	address, err := model.NeeAddressFromJsonBytes(bytes)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 2)
		if customer, err := customerUc.UpdateAddress(id, address); err != nil {
			render.JSON(w, r, err)
			return
		} else {
			render.JSON(w, r, customer)
		}
	} else {
		render.JSON(w, r, err)
	}
}

// TODO: implement DELETE CUSTOMER
func deleteCustomer(w http.ResponseWriter, r *http.Request) {

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := customerUc.Find(id); err == nil {
			render.JSON(w, r, customer)
			return
		} else {
			render.JSON(w, r, err)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func getCustomersPaginated(w http.ResponseWriter, r *http.Request) {
	if params, page, pageSize, err := getPaginatedParamsFromQuery(r.URL.Query()); err == nil {
		if customerUs, err := provider.CustomerUseCase(); err == nil {
			if customers, total, err := customerUs.List(params, page, pageSize); err != nil {
				render.JSON(w, r, err)
				return
			} else {
				render.JSON(w, r, newPaginatedResponse(customers, page, len(customers), total))
			}
		} else {
			render.JSON(w, r, err)
		}
	}
}
