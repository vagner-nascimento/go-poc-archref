package app

func setCustomerCreditCardHash(c *Customer) error {
	setCreditCardHash(c)
	if c.CreditCardHash == "" {
		return simpleError("cannot create customer's credit card hash")
	}

	return nil
}

func getCustomer(data []byte) (Customer, error) {
	c, err := makeCustomerFromBytes(data)
	if err != nil {
		return Customer{}, err
	}

	return c, nil
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
		return validationError(msgs)
	}

	return nil
}

func mergeUserToCustomer(u user, c Customer) Customer {
	return mapUserToCustomer(u, c)
}

func mergeCustomerUpdate(oldCustomer Customer, newCustomer Customer) Customer {
	return makeCustomerToUpdate(oldCustomer, newCustomer)
}
