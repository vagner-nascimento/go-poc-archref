package app

func CreateCustomer(c *Customer) error {
	return addCustomer(c)
}

func CreateUser(u *User) error {
	return nil
}
