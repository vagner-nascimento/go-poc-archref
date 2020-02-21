package app

func CreateCustomer(c *Customer, repository CustomerDataHandler) error {
	err := setCustomerCreditCardHash(c)
	if err == nil {
		err = repository.Save(c)
	}

	return err
}

func UpdateCustomer(c *Customer, repository CustomerDataHandler) error {
	if err := repository.Update(c); err != nil {
		return simpleError("cannot update the customer")
	}

	return nil
}
