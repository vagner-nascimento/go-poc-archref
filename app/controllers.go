package app

func AddCustomer(c *Customer, repository CustomerDataHandler) error {
	err := setCustomerCreditCardHash(c)
	if err == nil {
		err = repository.Save(c)
	}

	return err
}

func AddUser(u *User, repository UserDataHandler) error {
	err := validateUser(*u)
	if err == nil {
		err = repository.Save(u)
	}

	return err
}
