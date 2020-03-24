package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"strings"
)

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

func validateUser(u model.User) error {
	var msgs []string
	if u.Id == "" {
		msgs = append(msgs, "model.User id is required")
	}
	if u.Name == "" {
		msgs = append(msgs, "model.User name is required")
	}
	if u.EMail == "" {
		msgs = append(msgs, "model.User email is required")
	}
	if len(msgs) > 0 {
		return errors.New(strings.Join(msgs, ","))
	}
	return nil
}
