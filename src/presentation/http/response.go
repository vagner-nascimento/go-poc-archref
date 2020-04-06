package httppresentation

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"net/http"
	"reflect"
)

type paginatedResponse struct {
	Data     interface{} `json:"data"`
	Page     int64       `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
}

func newPaginatedResponse(data interface{}, page int64, quantity int, total int64) paginatedResponse {
	if reflect.ValueOf(data).IsNil() {
		data = make([]interface{}, 0)
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	return paginatedResponse{
		Data:     data,
		Page:     page,
		PageSize: quantity,
		Total:    total,
	}
}

func writeOkResponse(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func writeCreatedResponse(w http.ResponseWriter, data interface{}) {
	jsonData, _ := json.Marshal(data)

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func writeBadRequestResponse(w http.ResponseWriter, httpErr httpErrors) {
	jsonErr, _ := json.Marshal(httpErr)

	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonErr)
}

func writeInternalServerErrorResponse(w http.ResponseWriter, err error) {
	logger.Error("error http request", err)

	jsonErr, _ := json.Marshal(newInternalServerError())

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonErr)
}
