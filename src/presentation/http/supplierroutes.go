package httppresentation

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"io"
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
	supplier, vErr := parseAndValidateSupplier(r)
	if len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if supplierUc, err := provider.SupplierUseCase(); err == nil {
		if err = supplierUc.Create(&supplier); err != nil {
			render.JSON(w, r, err)
		} else {
			writeCreatedResponse(w, supplier)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func putSupplier(w http.ResponseWriter, r *http.Request) {
	supplier, vErr := parseAndValidateSupplier(r)
	if len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if supplierUc, err := provider.SupplierUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if sup, err := supplierUc.Update(id, supplier); err != nil {
			render.JSON(w, r, err)
			return
		} else {
			writeOkResponse(w, sup)
		}
	} else {
		render.JSON(w, r, err)
	}
}

func getSupplier(w http.ResponseWriter, r *http.Request) {
	if supUc, err := provider.SupplierUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if sup, err := supUc.Find(id); err == nil {
			writeOkResponse(w, sup)
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
				writeOkResponse(w, newPaginatedResponse(suppliers, page, len(suppliers), total))
			}
		} else {
			render.JSON(w, r, err)
		}
	}
}

func parseAndValidateSupplier(r *http.Request) (supplier model.Supplier, httpErr httpErrors) {
	supplier, err := parseSupplierFromBody(r.Body)
	if err != nil {
		httpErr.Errors = append(httpErr.Errors, newConversionError(err))
	} else {
		httpErr = newValidationErrors(supplier.Validate())
	}

	return supplier, httpErr
}

func parseSupplierFromBody(body io.ReadCloser) (supplier model.Supplier, err error) {
	defer body.Close()
	if bytes, err := ioutil.ReadAll(body); err == nil {
		supplier, err = model.NewSupplierFromJsonBytes(bytes)
	}

	return supplier, err
}
