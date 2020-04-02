package model

import "encoding/json"

// TODO: realise how to include dates into models
type Customer struct {
	Id         string  `json:"id" bson:"id"`
	Name       string  `json:"name" validate:"required,min=3,max=150" bson:"name"`
	EMail      string  `json:"eMail" validate:"required,email" bson:"eMail"`
	BirthYear  int     `json:"birthYear" validate:"min=1900" bson:"birthYear"`
	BirthDay   int     `json:"birthDay" validate:"min=1,max=31" bson:"birthDay"`
	BirthMonth int     `json:"birthMonth" validate:"min=1,max=12" bson:"birthMont"`
	UserId     string  `json:"userId" bson:"userId"`
	Address    Address `json:"address" validate:"required" bson:"address"`
}

type Address struct {
	Street       string `json:"street" validate:"required,min=3,max=255" bson:"street"`
	Number       string `json:"number" validate:"required,min=2,max=255" bson:"number"`
	Neighborhood string `json:"neighborhood" validate:"required,min=3,max=255" bson:"neighborhood"`
	Complement   string `json:"complement" validate:"min=2,max=255" bson:"complement"`
	PostalCode   string `json:"postalCode" validate:"required,min=2,max=150" bson:"postalCode"`
	City         string `json:"city" validate:"required,min=3,max=150" bson:"city"`
	Country      string `json:"country" validate:"required,min=2,max=2" bson:"country"`
	State        string `json:"state" validate:"required,min=2,max=150" bson:"state"`
}

func NewCustomerFromJsonBytes(bytes []byte) (customer Customer, err error) {
	err = json.Unmarshal(bytes, &customer)
	return customer, err
}

func NewAddressFromJsonBytes(data []byte) (address Address, err error) {
	err = json.Unmarshal(data, &address)
	return address, err
}
