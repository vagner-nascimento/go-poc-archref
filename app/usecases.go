package app

import "errors"

func createCustomer(c *Customer) *Customer {
	c.setCreditCardHash()

	return c
}

func validateUser(u User) error {
	if u.CustomerId == "" {
		return errors.New("customer id not informed")
	}

	return nil
}
