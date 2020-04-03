package model

import "encoding/json"

// TODO: realise how to include dates into models
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

func NewCustomerFromJsonBytes(bytes []byte) (customer Customer, err error) {
	err = json.Unmarshal(bytes, &customer)
	return customer, err
}

func NewAddressFromJsonBytes(data []byte) (address Address, err error) {
	err = json.Unmarshal(data, &address)
	return address, err
}
