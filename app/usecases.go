package app

func setCustomerCreditCardHash(c *Customer) error {
	setCreditCardHash(c)
	if c.CreditCardHash == "" {
		return simpleError("cannot create customer's credit card hash")
	}

	return nil
}

func getCustomer(data interface{}) (Customer, error) {
	c, err := makeCustomer(data)
	if err != nil {
		return Customer{}, err
	}

	return c, nil
}

func validateUpdateData(c Customer) error {
	var msgs []string
	if c.Id == "" && c.UserId == "" {
		msgs = append(msgs, "neither id nor user id are informed")
	}
	if c.Name == "" {
		msgs = append(msgs, "name is not informed")
	}
	if c.EMail == "" {
		msgs = append(msgs, "email is not informed")
	}

	if len(msgs) > 0 {
		return validationError(msgs)
	}

	return nil
}
