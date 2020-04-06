package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

//TODO: Realise how to make invalid when bool is not informed (pointer is an way)
type Supplier struct {
	Id             string  `json:"id" bson:"id"`
	Name           string  `json:"name" bson:"name" validate:"required,min=3,max=150"`
	DocumentNumber string  `json:"documentNumber" bson:"documentNumber" validate:"required,min=3,max=150"`
	IsActive       bool    `json:"isActive" bson:"isActive"`
	CreditLimit    float64 `json:"creditLimit" bson:"creditLimit" validate:"min=1"`
}

func (s *Supplier) Validate() (valErrs validator.ValidationErrors) {
	v := validator.New()
	if err := v.Struct(*s); err != nil {
		valErrs = err.(validator.ValidationErrors)
	}

	return valErrs
}

func NewSupplierFromJsonBytes(data []byte) (sup Supplier, err error) {
	err = json.Unmarshal(data, &sup)
	return sup, err
}
