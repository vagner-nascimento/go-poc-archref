package presentation

import (
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