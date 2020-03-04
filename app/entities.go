package app

import (
	"encoding/json"
)

type Customer struct {
	Id         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	EMail      string `json:"eMail" bson:"eMail"`
	BirthYear  int    `json:"birthYear" bson:"birthYear"`
	BirthDay   int    `json:"birthDay" bson:"birthDay"`
	BirthMonth int    `json:"birthMonth" bson:"birthMont"`
	UserId     string `json:"userId" bson:"userId"`
}

type user struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	EMail      string `json:"eMail"`
	BirthYear  int    `json:"birthYear"`
	BirthDay   int    `json:"birthDay"`
	BirthMonth int    `json:"birthMonth"`
}

type SearchParameter struct {
	Field string
	Value string
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
		Id:         c.Id,
		Name:       u.Name,
		EMail:      u.EMail,
		BirthYear:  u.BirthYear,
		BirthDay:   u.BirthDay,
		BirthMonth: u.BirthMonth,
		UserId:     u.Id,
	}
}

func makeCustomerToUpdate(oldCustomer Customer, newCustomer Customer) Customer {
	return Customer{
		Id:         oldCustomer.Id,
		Name:       newCustomer.Name,
		EMail:      newCustomer.EMail,
		BirthYear:  newCustomer.BirthYear,
		BirthDay:   newCustomer.BirthDay,
		BirthMonth: newCustomer.BirthMonth,
		UserId:     newCustomer.UserId,
	}
}
