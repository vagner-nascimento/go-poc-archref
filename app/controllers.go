package app

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

	customers, err := repository.GetMany([]SearchParameter{{
		Field:    "email",
		Operator: "=",
		Value:    u.EMail,
	}})

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

	// TODO: finish update
	//if err = repository.Update(&newCustomer); err != nil {
	//	return Customer{}, err
	//}

	return newCustomer, nil
}

func FindCustomer(id string, repository CustomerDataHandler) (Customer, error) {
	return repository.Get(id)
}
