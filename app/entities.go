package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Customer struct {
	Id             string `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	EMail          string `json:"eMail" bson:"eMail"`
	CreditCardHash string `json:"-" bson:"creditCardHash"`
	BirthYear      int    `json:"birthYear" bson:"birthYear"`
	BirthDay       int    `json:"birthDay" bson:"birthDay"`
	BirthMonth     int    `json:"birthMonth" bson:"birthMont"`
	UserId         string `json:"userId" bson:"userId"`
}

type SearchParameter struct {
	Field    string
	Operator string
	Value    string
}

type UpdateParameter struct {
	Field string
	Value string
}

type user struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	EMail      string `json:"eMail"`
	BirthYear  int    `json:"birthYear"`
	BirthDay   int    `json:"birthDay"`
	BirthMonth int    `json:"birthMonth"`
}

func makeCustomerFromBytes(bytes []byte) (Customer, error) {
	var c Customer
	if err := json.Unmarshal(bytes, &c); err != nil {
		return c, conversionError("bytes", "customer")
	}

	return c, nil
}

func makeUserFromBytes(data []byte) (user, error) {
	var u user
	if err := json.Unmarshal(data, &u); err != nil {
		return u, conversionError("bytes", "user")
	}

	return u, nil
}

func mapUserToCustomer(u user, c Customer) Customer {
	return Customer{
		Id:             c.Id,
		Name:           u.Name,
		EMail:          u.EMail,
		CreditCardHash: c.CreditCardHash,
		BirthYear:      u.BirthYear,
		BirthDay:       u.BirthDay,
		BirthMonth:     u.BirthMonth,
		UserId:         u.Id,
	}
}

func makeCustomerToUpdate(oldCustomer Customer, newCustomer Customer) Customer {
	return Customer{
		Id:             oldCustomer.Id,
		Name:           newCustomer.Name,
		EMail:          newCustomer.EMail,
		CreditCardHash: oldCustomer.CreditCardHash,
		BirthYear:      newCustomer.BirthYear,
		BirthDay:       newCustomer.BirthDay,
		BirthMonth:     newCustomer.BirthMonth,
		UserId:         newCustomer.UserId,
	}
}

func setCreditCardHash(c *Customer) {
	c.CreditCardHash = "fake_"
	for i := 0; i < 5; i = i + 1 {
		c.CreditCardHash += strings.Split(fmt.Sprintf("%f", rand.Float64()), ".")[1]
	}
}
