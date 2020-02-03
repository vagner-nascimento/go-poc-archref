package app

import (
	"encoding/json"
)

type person struct {
	Id   string `id: "id"`
	Name string `name: "name"`
}

type customer struct {
	person
	Alias          string `alias: "alias"`
	CreditCardHash string
}

func makeCustomer(data []byte) customer {
	var c customer
	json.Unmarshal(data, &c)

	return c
}
