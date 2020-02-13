package app

func AddCustomer(c *Customer, data CustomerDataHandler) error {
	return data.Save(createCustomer(c))
}

func AddUser(u *User, data UserDataHandler) error {
	err := validateUser(*u)
	if err == nil {
		return data.Save(u)
	}

	return err
}
