package model

import "encoding/json"

type Supplier struct {
	Id             string  `json:"id" bson:"id"`
	Name           string  `json:"name" bson:"name" validate:"required,min=3,max=150"`
	DocumentNumber string  `json:"documentNumber" bson:"documentNumber" validate:"required,min=3,max=150"`
	IsActive       bool    `json:"isActive" bson:"isActive" validate:"required"`
	CreditLimit    float64 `json:"creditLimit" bson:"creditLimit" validate:"min=1"`
}

func NewSupplierFromJsonBytes(data []byte) (sup Supplier, err error) {
	err = json.Unmarshal(data, &sup)
	return sup, err
}
