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

type User struct {
	Id         string `id: "id"`
	Name       string `name: "name"`
	EMail      string `eMail: "eMail"`
	BirthYear  int    `birthYear: "birthYear"`
	BirthDay   int    `birthDay: "birthDay"`
	BirthMonth int    `birthMont: "birthMonth"`
}

func setCreditCardHash(c *Customer) {
	c.CreditCardHash = "fake_"
	for i := 0; i < 5; i = i + 1 {
		c.CreditCardHash += strings.Split(fmt.Sprintf("%f", rand.Float64()), ".")[1]
	}
}

func MakeUserFromBytes(data []byte) (User, error) {
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		return u, conversionError("bytes", "user")
	}

	return u, nil
}

func makeCustomer(data interface{}) (Customer, error) {
	switch t := data.(type) {
	case User:
		return buildCustomerFromUser(data.(User)), nil
	case *User:
		return buildCustomerFromUser(*data.(*User)), nil
	case []byte:
		return buildCustomerFromBytes(data.([]byte))
	case Customer:
		return data.(Customer), nil
	default:
		return Customer{}, simpleError(fmt.Sprintf("invalid data type %s to build a customer", t))
	}
}

func buildCustomerFromBytes(bytes []byte) (Customer, error) {
	var c Customer
	if err := json.Unmarshal(bytes, &c); err != nil {
		return c, conversionError("bytes", "customer")
	}

	return c, nil
}

func buildCustomerFromUser(u User) Customer {
	return Customer{
		Id:             "",
		Name:           u.Name,
		EMail:          u.EMail,
		CreditCardHash: "",
		BirthYear:      u.BirthDay,
		BirthDay:       u.BirthDay,
		BirthMonth:     u.BirthYear,
		UserId:         u.Id,
	}
}
