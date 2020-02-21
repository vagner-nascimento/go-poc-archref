package app

func setCustomerCreditCardHash(c *Customer) error {
	c.setCreditCardHash()

	if c.CreditCardHash == "" {
		return simpleError("cannot create customer's credit card hash")
	}

	return nil
}
