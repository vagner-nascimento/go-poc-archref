package app

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Customer struct {
	Id             string `id: "id"`
	Name           string `name: "name"`
	EMail          string `eMail: "eMail"`
	CreditCardHash string
	BirthYear      int `birthYear: "birthYear"`
	BirthDay       int `birthDay: "birthDay"`
	BirthMonth     int `birthMont: "birthMonth"`
	UserId         string
}

type SearchParameter struct {
	Field    string
	Operator string
	Value    string
}

type user struct {
	Id         string `id: "id"`
	Name       string `name: "name"`
	EMail      string `eMail: "eMail"`
	BirthYear  int    `birthYear: "birthYear"`
	BirthDay   int    `birthDay: "birthDay"`
	BirthMonth int    `birthMont: "birthMonth"`
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

func setCreditCardHash(c *Customer) {
	c.CreditCardHash = "fake_"
	for i := 0; i < 5; i = i + 1 {
		c.CreditCardHash += strings.Split(fmt.Sprintf("%f", rand.Float64()), ".")[1]
	}
}
