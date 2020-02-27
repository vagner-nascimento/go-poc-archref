package app

import "fmt"

func CreateCustomer(data []byte, repository CustomerDataHandler) (interface{}, error) {
	c, err := getCustomer(data)
	if err == nil {
		if err = setCustomerCreditCardHash(&c); err == nil {
			err = repository.Save(&c)
		}
	}

	return c, err
}

func UpdateCustomerFromUser(data []byte, repository CustomerDataHandler) (Customer, error) {
	u, err := getUser(data)
	if err != nil {
		return Customer{}, err
	}

	if err = validateUser(u); err != nil {
		return Customer{}, err
	}

	customers, err := repository.GetMany(fmt.Sprintf("eMail: %s", u.EMail))
	if err != nil {
		return Customer{}, err
	}

	var foundCustomer Customer
	if len(customers) > 0 {
		foundCustomer = customers[0]
	}

	if foundCustomer.Id == "" {
		return Customer{}, notFoundError("customer")
	}

	newCustomer := mergeUserToCustomer(u, foundCustomer)
	if err = repository.Update(&newCustomer); err != nil {
		return Customer{}, err
	}

	return newCustomer, nil
}
