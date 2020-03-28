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
