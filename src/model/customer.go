package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
)

// TODO: realise how to include dates into models
// TODO: realise how to save and return null when none data
type Customer struct {
	Id         string  `json:"id" bson:"id"`
	Name       string  `json:"name" bson:"name" validate:"required,min=3,max=150"`
	EMail      string  `json:"eMail" bson:"eMail" validate:"required,email"`
	BirthYear  int     `json:"birthYear" bson:"birthYear" validate:"min=1900"`
	BirthDay   int     `json:"birthDay" bson:"birthDay" validate:"min=1,max=31"`
	BirthMonth int     `json:"birthMonth" bson:"birthMont" validate:"min=1,max=12"`
	UserId     string  `json:"userId" bson:"userId"`
	Address    Address `json:"address" bson:"address" validate:"required"`
}

// TODO: validate on AMQP consumers or into app Use Cases or Entity
// TODO: realise how to validate optionals only when informed (if not informed or null should be valid)
func (c *Customer) Validate() (valErrs validator.ValidationErrors) {
	v := validator.New()
	if err := v.Struct(*c); err != nil {
		valErrs = err.(validator.ValidationErrors)
	}

	return valErrs
}

func NewCustomerFromJsonBytes(bytes []byte) (customer Customer, err error) {
	err = json.Unmarshal(bytes, &customer)
	return customer, err
}
