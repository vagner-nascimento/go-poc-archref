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
}

func (c *Customer) setCreditCardHash() {
	c.CreditCardHash = "fake_"
	for i := 0; i < 5; i = i + 1 {
		c.CreditCardHash += strings.Split(fmt.Sprintf("%f", rand.Float64()), ".")[1]
	}
}

type User struct {
	CustomerId string `customerId: "customerId"`
	Name       string `name: "name"`
	UserName   string `userName: "userName"`
	EMail      string `eMail: "eMail"`
}

func BuildCustomerFromBytes(data []byte) (Customer, error) {
	var c Customer

	err := json.Unmarshal(data, &c)
	if err != nil {
		return c, conversionError(err, "create a new customer form bytes")
	}

	return c, nil
}

func BuildUserFromCustomer(c Customer) User {
	return User{
		CustomerId: c.Id,
		Name:       c.Name,
		UserName:   strings.Split(c.Name, " ")[0],
		EMail:      c.EMail,
	}
}
