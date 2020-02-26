package app

import "fmt"

func CreateCustomer(c *Customer, repository CustomerDataHandler) error {
	err := setCustomerCreditCardHash(c)
	if err == nil {
		err = repository.Save(c)
	}

	return err
}

func UpdateCustomer(data interface{}, repository CustomerDataHandler) (Customer, error) {
	c, err := getCustomer(data)
	if err != nil {
		return Customer{}, err
	}

	if err = validateUpdateData(c); err != nil {
		return Customer{}, err
	}

	var foundCustomer Customer
	if c.Id != "" {
		foundCustomer, err = repository.Get(c.Id)

		if err != nil {
			return Customer{}, err
		}
	} else {
		customers, err := repository.GetMany(fmt.Sprintf("userId: %s", c.UserId))
		if err != nil {
			return Customer{}, err
		}

		if len(customers) > 0 {
			foundCustomer = customers[0]
		}
	}

	if foundCustomer.Id == "" {
		return Customer{}, notFoundError("customer")
	}

	//TODO: Make merge of data and update it

	return c, nil
}
