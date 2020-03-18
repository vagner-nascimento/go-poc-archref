package app

import (
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

// TODO: think in better way to pass parameters to the queries
type SearchParameter struct {
	Field  string
	Values []interface{}
}

func mapUserToCustomer(u model.User, c model.Customer) model.Customer {
	return model.Customer{
		Id:         c.Id,
		Name:       u.Name,
		EMail:      u.EMail,
		BirthYear:  u.BirthYear,
		BirthDay:   u.BirthDay,
		BirthMonth: u.BirthMonth,
		UserId:     u.Id,
	}
}

func mapCustomerToUpdate(oldCustomer model.Customer, newCustomer model.Customer) model.Customer {
	return model.Customer{
		Id:         oldCustomer.Id,
		Name:       newCustomer.Name,
		EMail:      newCustomer.EMail,
		BirthYear:  newCustomer.BirthYear,
		BirthDay:   newCustomer.BirthDay,
		BirthMonth: newCustomer.BirthMonth,
		UserId:     newCustomer.UserId,
	}
}
