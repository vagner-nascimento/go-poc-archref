package app

import (
	"errors"
	"strings"
)

func getCustomer(data []byte) (Customer, error) {
	c, err := makeCustomerFromBytes(data)
	if err != nil {
		return Customer{}, err
	}

	return c, nil
}

func getCustomerToUpdate(oldCustomer Customer, data []byte) (newCustomer Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = errors.New("customer not found")
		return newCustomer, err
	}

	newData, err := makeCustomerFromBytes(data)
	if err != nil {
		return newCustomer, err
	}

	newCustomer = mapCustomerToUpdate(oldCustomer, newData)
	return newCustomer, err
}

func getCustomerToUpdateEmail(oldCustomer Customer, data []byte) (newCustomer Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = errors.New("customer not found")
		return newCustomer, err
	}

	newData, err := makeCustomerFromBytes(data)
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

func getUser(data []byte) (user, error) {
	u, err := makeUserFromBytes(data)
	if err != nil {
		return user{}, err
	}

	return u, nil
}

func validateUser(u user) error {
	var msgs []string
	if u.Id == "" {
		msgs = append(msgs, "user id is required")
	}

	if u.Name == "" {
		msgs = append(msgs, "user name is required")
	}

	if u.EMail == "" {
		msgs = append(msgs, "user email is required")
	}

	if len(msgs) > 0 {
		return errors.New(strings.Join(msgs, ","))
	}

	return nil
}

func mergeUserToCustomer(u user, c Customer) Customer {
	return mapUserToCustomer(u, c)
}
