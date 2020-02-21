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

func BuildCustomerFromBytes(data []byte) (Customer, error) {
	var c Customer

	err := json.Unmarshal(data, &c)
	if err != nil {
		return c, conversionError(err, "bytes", "customer")
	}

	return c, nil
}
