package app

func setCustomerCreditCardHash(c *Customer) error {
	setCreditCardHash(c)
	if c.CreditCardHash == "" {
		return simpleError("cannot create customer's credit card hash")
	}

	return nil
}

func getCustomer(data interface{}) (Customer, error) {
	return Customer{}, nil
}
