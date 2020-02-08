package app

func CreateCustomer(c *Customer) error {
	err := addCustomer(c)
	if err != nil {
		return err
	}

	return nil
}
