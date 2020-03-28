package model

import "encoding/json"

type Supplier struct {
	Id             string  `json:"id" bson:"id"`
	Name           string  `json:"name" bson:"name"`
	DocumentNumber string  `json:"documentNumber" bson:"documentNumber"`
	CreditLimit    float64 `json:"creditLimit" bson:"creditLimit"`
}

func NewSupplierFromJsonBytes(data []byte) (sup Supplier, err error) {
	err = json.Unmarshal(data, &sup)
	return sup, err
}
