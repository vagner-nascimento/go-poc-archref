package app

func setCustomerCreditCardHash(c *Customer) error {
	c.setCreditCardHash()

	if c.CreditCardHash == "" {
		return simpleError("cannot create customer's credit card hash")
	}

	return nil
}

func validateUser(u User) error {
	if u.CustomerId == "" {
		return simpleError("cannot create customer's card hash")
	}

	return nil
}
