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

type User struct {
	person
	UseName    string `userName: "userName"`
	Password   string
	repository UserDataIterable
}

func NewCustomer(repo CustomerDataIterable) (Customer, error) {
	var c Customer
	if repo == nil {
		return c, errors.New("Repository must be informed")
	}

	c = Customer{repository: repo}

	return c, nil
}

func NewUser(repo UserDataIterable) (User, error) {
	var c User
	if repo == nil {
		return c, errors.New("Repository must be informed")
	}

	c = User{repository: repo}

	return c, nil
}
