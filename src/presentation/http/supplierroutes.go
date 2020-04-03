package presentation

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"io/ioutil"
	"net/http"
)

func newSupplierRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", postSupplier)
	router.Put("/{id}", putSupplier)
	router.Get("/{id}", getSupplier)
	router.Get("/", getSuppliersPaginated)
	return router
}

func postSupplier(w http.ResponseWriter, r *http.Request) {
	supplier, err := parseSupplier(r)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if vErr := validateBody(supplier); len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if supplierUc, err := provider.SupplierUseCase(); err == nil {
		if err = supplierUc.Create(&supplier); err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, supplier)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func putSupplier(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	sup, err := model.NewSupplierFromJsonBytes(bytes)
	if err != nil {
		render.JSON(w, r, err)
		return
	}
	if supplierUc, err := provider.SupplierUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := supplierUc.Update(id, sup); err != nil {
			render.JSON(w, r, err)
			return
		} else {
			render.JSON(w, r, customer)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func getSupplier(w http.ResponseWriter, r *http.Request) {
	if supUc, err := provider.SupplierUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := supUc.Find(id); err == nil {
			render.JSON(w, r, customer)
			return
		} else {
			render.JSON(w, r, err)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func getSuppliersPaginated(w http.ResponseWriter, r *http.Request) {
	if params, page, pageSize, err := getPaginatedParamsFromQuery(r.URL.Query()); err == nil {
		if supUc, err := provider.SupplierUseCase(); err == nil {
			if suppliers, total, err := supUc.List(params, page, pageSize); err != nil {
				render.JSON(w, r, err)
				return
			} else {
				render.JSON(w, r, newPaginatedResponse(suppliers, page, len(suppliers), total))
			}
		} else {
			render.JSON(w, r, err)
		}
	}
}

func parseSupplier(r *http.Request) (supplier model.Supplier, err error) {
	defer r.Body.Close()

	bytes, err := ioutil.ReadAll(r.Body)
	if err == nil {
		//If int(and i guess that other number too) starts with zero returns error
		if supplier, err = model.NewSupplierFromJsonBytes(bytes); err != nil {
			supplier = model.Supplier{}
		}
	}

	return supplier, err
}
