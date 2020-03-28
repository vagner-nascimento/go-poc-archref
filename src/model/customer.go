package model

import "encoding/json"

type Customer struct {
	Id         string  `json:"id" bson:"id"`
	Name       string  `json:"name" bson:"name"`
	EMail      string  `json:"eMail" bson:"eMail"`
	BirthYear  int     `json:"birthYear" bson:"birthYear"`
	BirthDay   int     `json:"birthDay" bson:"birthDay"`
	BirthMonth int     `json:"birthMonth" bson:"birthMont"`
	UserId     string  `json:"userId" bson:"userId"`
	Address    Address `json:"address" bson:"address"`
}

func NewCustomerFromJsonBytes(bytes []byte) (customer Customer, err error) {
	err = json.Unmarshal(bytes, &customer)
	return customer, err
}

type Address struct {
	Street       string `json:"street" bson:"street"`
	Number       string `json:"number" bson:"number"`
	Neighborhood string `json:"neighborhood" bson:"neighborhood"`
	PostalCode   string `json:"postalCode" bson:"postalCode"`
	City         string `json:"city" bson:"city"`
	Country      string `json:"country" bson:"country"`
	State        string `json:"state" bson:"state"`
}

func NeeAddressFromJsonBytes(data []byte) (address Address, err error) {
	err = json.Unmarshal(data, &address)
	return address, err
}
