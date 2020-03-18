package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"strings"
)

func getCustomer(data []byte) (model.Customer, error) {
	return model.NewCustomerFromJsonBytes(data)
}

func getCustomerToUpdate(oldCustomer model.Customer, data []byte) (newCustomer model.Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = errors.New("customer not found")
		return newCustomer, err
	}

	newData, err := model.NewCustomerFromJsonBytes(data)
	if err != nil {
		return newCustomer, err
	}

	newCustomer = mapCustomerToUpdate(oldCustomer, newData)
	return newCustomer, err
}

func getCustomerToUpdateEmail(oldCustomer model.Customer, data []byte) (newCustomer model.Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = errors.New("customer not found")
		return newCustomer, err
	}

	newData, err := model.NewCustomerFromJsonBytes(data)
	if err != nil {
		return newCustomer, err
	}
	if newData.EMail == "" {
		return newCustomer, errors.New("email must be informed")
	}

	newCustomer = oldCustomer
	newCustomer.EMail = newData.EMail

	return newCustomer, err
}

func getUser(data []byte) (model.User, error) {
	return model.NewUserFromJsonBytes(data)
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

func mergeUserToCustomer(u model.User, c model.Customer) model.Customer {
	return mapUserToCustomer(u, c)
}
