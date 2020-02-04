package app

import (
	"encoding/json"
)

type person struct {
	Id   string `id: "id"`
	Name string `name: "name"`
}

type Customer struct {
	person
	Alias          string `alias: "alias"`
	CreditCardHash string
}

func makeCustomer(data []byte) Customer {
	var c Customer
	json.Unmarshal(data, &c)

	return c
}
