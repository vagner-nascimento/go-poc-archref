package app

import "errors"

type person struct {
	Id   string `id: "id"`
	Name string `name: "name"`
}

type Customer struct {
	person
	Alias          string `alias: "alias"`
	CreditCardHash string
	repository     CustomerDataIterable
}

func NewCustomer(repo CustomerDataIterable) (Customer, error) {
	var c Customer
	if repo == nil {
		return c, errors.New("Repository must be informed")
	}

	c = Customer{repository: repo}

	return c, nil
}
