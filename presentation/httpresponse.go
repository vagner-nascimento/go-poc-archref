package presentation

type paginatedResponse struct {
	data     interface{} `json:"data"`
	page     int64       `json:"page"`
	quantity int         `json:"quantity"`
	total    int64       `json:"total"`
}

func newPaginatedResponse(data interface{}, page int64, quantity int, total int64) paginatedResponse {
	return paginatedResponse{
		data:     data,
		page:     page,
		quantity: quantity,
		total:    total,
	}
}
