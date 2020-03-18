package app

func getCustomer(data []byte) (Customer, error) {
	c, err := makeCustomerFromBytes(data)
	if err != nil {
		return Customer{}, err
	}

	return c, nil
}

func getCustomerToUpdate(oldCustomer Customer, data []byte) (newCustomer Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = customerNotFoundError()
		return
	}

	newData, err := makeCustomerFromBytes(data)
	if err != nil {
		return
	}

	newCustomer = mapCustomerToUpdate(oldCustomer, newData)
	return newCustomer, err
}

func getCustomerToUpdateEmail(oldCustomer Customer, data []byte) (newCustomer Customer, err error) {
	if len(oldCustomer.Id) <= 0 {
		err = customerNotFoundError()
		return
	}

	newData, err := makeCustomerFromBytes(data)
	if err != nil {
		return
	}
	if newData.EMail == "" {
		return newCustomer, validationError([]string{"email must be informed"})
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
		return validationError(msgs)
	}

	return nil
}

func mergeUserToCustomer(u user, c Customer) Customer {
	return mapUserToCustomer(u, c)
}
