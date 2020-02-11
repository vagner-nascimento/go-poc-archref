package app

import (
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
	data           CustomerDataHandler
}

func (c *Customer) save() error {
	return c.data.Save(c)
}

type User struct {
	CustomerId string `customerId: "customerId"`
	Name       string `name: "name"`
	UserName   string `userName: "userName"`
	EMail      string `eMail: "eMail"`
}

func NewCustomer(db CustomerDataHandler) *Customer {
	return &Customer{data: db}
}

func NewUserFromCustomer(c Customer) User {
	return mapCustomerIntoUser(c)
}

func mapCustomerIntoUser(c Customer) User {
	return User{
		CustomerId: c.Id,
		Name:       c.Name,
		UserName:   strings.Split(c.Name, " ")[0],
		EMail:      c.EMail,
	}
}
