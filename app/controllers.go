package app

func CreateCustomer(c *Customer, repository CustomerDataHandler) error {
	err := setCustomerCreditCardHash(c)
	if err == nil {
		err = repository.Save(c)
	}

	return err
}

func UpdateCustomer(data interface{}, repository CustomerDataHandler) (*Customer, error) {
	c, err := getCustomer(data)
	if err != nil {
		return &c, err
	}
	return &c, nil
}
