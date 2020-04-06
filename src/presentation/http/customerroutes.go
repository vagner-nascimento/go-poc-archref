package httppresentation

import (
	"github.com/go-chi/chi"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
	"io"
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

// TODO: implement DELETE CUSTOMER
func deleteCustomer(w http.ResponseWriter, r *http.Request) {

}

// TODO: validate params (path, query)
// TODO: clean duplicate codes
func postCustomer(w http.ResponseWriter, r *http.Request) {
	customer, vErr := parseAndValidateCustomer(r)
	if len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		if err = customerUc.Create(&customer); err != nil {
			writeInternalServerErrorResponse(w, err)
		} else {
			writeCreatedResponse(w, customer)
		}
	} else {
		writeInternalServerErrorResponse(w, err)
	}
}

// TODO: write not found response
func putCustomer(w http.ResponseWriter, r *http.Request) {
	customer, vErr := parseAndValidateCustomer(r)
	if len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := customerUc.Update(id, customer); err != nil {
			writeInternalServerErrorResponse(w, err)
		} else {
			writeOkResponse(w, customer)
		}
	} else {
		writeInternalServerErrorResponse(w, err)
	}
}

func patchCustomerAddress(w http.ResponseWriter, r *http.Request) {
	address, vErr := parseAndValidateAddress(r)
	if len(vErr.Errors) > 0 {
		writeBadRequestResponse(w, vErr)
		return
	}
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 2)
		if customer, err := customerUc.UpdateAddress(id, address); err != nil {
			writeInternalServerErrorResponse(w, err)
		} else {
			writeOkResponse(w, customer)
		}
	} else {
		writeInternalServerErrorResponse(w, err)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	if customerUc, err := provider.CustomerUseCase(); err == nil {
		id := getDataFromPath(r.URL.Path, 1)
		if customer, err := customerUc.Find(id); err == nil {
			writeOkResponse(w, customer)
		} else {
			writeInternalServerErrorResponse(w, err)
		}
	} else {
		writeInternalServerErrorResponse(w, err)
	}
}

func getCustomersPaginated(w http.ResponseWriter, r *http.Request) {
	if params, page, pageSize, err := getPaginatedParamsFromQuery(r.URL.Query()); err == nil {
		if customerUs, err := provider.CustomerUseCase(); err == nil {
			if customers, total, err := customerUs.List(params, page, pageSize); err != nil {
				writeInternalServerErrorResponse(w, err)
			} else {
				writeOkResponse(w, newPaginatedResponse(customers, page, len(customers), total))
			}
		} else {
			writeInternalServerErrorResponse(w, err)
		}
	} else {
		writeInternalServerErrorResponse(w, err)
	}
}

func parseAndValidateCustomer(r *http.Request) (customer model.Customer, httpErr httpErrors) {
	customer, err := parseCustomerFromBody(r.Body)
	if err != nil {
		httpErr.Errors = append(httpErr.Errors, newConversionError(err))
	} else {
		httpErr = newValidationErrors(customer.Validate())
	}

	return customer, httpErr
}

func parseCustomerFromBody(body io.ReadCloser) (customer model.Customer, err error) {
	defer body.Close()

	if bytes, err := ioutil.ReadAll(body); err == nil {
		customer, err = model.NewCustomerFromJsonBytes(bytes)
	}

	return customer, err
}

func parseAndValidateAddress(r *http.Request) (address model.Address, httpErr httpErrors) {
	address, err := parseAddressFromBody(r.Body)
	if err != nil {
		httpErr.Errors = append(httpErr.Errors, newConversionError(err))
	} else {
		httpErr = newValidationErrors(address.Validate())
	}

	return address, httpErr
}

func parseAddressFromBody(body io.ReadCloser) (address model.Address, err error) {
	defer body.Close()
	if bytes, err := ioutil.ReadAll(body); err == nil {
		address, err = model.NewAddressFromJsonBytes(bytes)
	}

	return address, err
}
