package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

type Address struct {
	Street       string `json:"street" bson:"street" validate:"required,min=3,max=255"`
	Number       string `json:"number" bson:"number" validate:"required,min=2,max=255"`
	Neighborhood string `json:"neighborhood" bson:"neighborhood" validate:"required,min=3,max=255"`
	Complement   string `json:"complement" bson:"complement" validate:"min=2,max=255"`
	PostalCode   string `json:"postalCode" bson:"postalCode" validate:"required,min=2,max=150"`
	City         string `json:"city" bson:"city" validate:"required,min=3,max=150"`
	Country      string `json:"country" bson:"country" validate:"required,min=2,max=2"`
	State        string `json:"state" bson:"state" validate:"required,min=2,max=150"`
}

func (a *Address) Validate() (valErrs validator.ValidationErrors) {
	v := validator.New()
	if err := v.Struct(*a); err != nil {
		valErrs = err.(validator.ValidationErrors)
	}

	return valErrs
}

func NewAddressFromJsonBytes(data []byte) (address Address, err error) {
	err = json.Unmarshal(data, &address)

	return address, err
}
