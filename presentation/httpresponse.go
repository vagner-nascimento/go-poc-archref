package presentation

type paginatedResponse struct {
	Data     interface{} `json:"data"`
	Page     int64       `json:"page"`
	Quantity int         `json:"quantity"`
	Total    int64       `json:"total"`
}

// TODO: realise how to receive only arrays in data
func newPaginatedResponse(data interface{}, page int64, quantity int, total int64) paginatedResponse {
	if page == 0 {
		page = 1
	}

	switch val := data.(type) {
	default:
		if val == nil { // TODO: even val been nil, it doesn't works
			var emptyArr []interface{}
			data = emptyArr
		}
	}

	return paginatedResponse{
		Data:     data,
		Page:     page,
		Quantity: quantity,
		Total:    total,
	}
}
