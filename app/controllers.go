package app

// TODO: review return without var names
/*
	search in project by
	return
}
*/
func CreateCustomer(customerData []byte, repository CustomerDataHandler) (Customer, error) {
	customer, err := getCustomer(customerData)
	if err == nil {
		if err = repository.Save(&customer); err != nil {
			customer = Customer{}
		}
	}

	return customer, err
}

func UpdateCustomer(id string, customerData []byte, repository CustomerDataHandler) (newCustomer Customer, err error) {
	var foundCustomer Customer
	foundCustomer, err = repository.Get(id)
	if err != nil {
		return
	}

	newCustomer, err = getCustomerToUpdate(foundCustomer, customerData)
	if err != nil {
		return
	}

	if err = repository.Replace(newCustomer); err != nil {
		newCustomer = Customer{}
		return
	}

	return newCustomer, err
}

func UpdateCustomerFromUser(userData []byte, repository CustomerDataHandler) (Customer, error) {
	user, err := getUser(userData)
	if err != nil {
		return Customer{}, err
	}

	if err = validateUser(user); err != nil {
		return Customer{}, err
	}

	// TODO: think in a better way to send params to the repo
	customers, total, err := repository.GetMany([]SearchParameter{{
		Field: "EMail",
		Value: user.EMail,
	}},
		0,
		2)

	if total > 1 {
		return Customer{}, validationError([]string{"to many register with the same e-mail"})
	} else if total == 0 {
		return Customer{}, customerNotFoundError()
	}

	newCustomer := mergeUserToCustomer(user, customers[0])
	if err = repository.Replace(newCustomer); err != nil {
		return Customer{}, err
	}

	return newCustomer, nil
}

func FindCustomer(id string, repository CustomerDataHandler) (Customer, error) {
	if customer, err := repository.Get(id); err != nil {
		return Customer{}, err
	} else if len(customer.Id) <= 0 {
		return Customer{}, customerNotFoundError()
	} else {
		return customer, nil
	}
}

func FindCustomers(params []SearchParameter, page int64, quantity int64, repository CustomerDataHandler) (res []Customer, total int64, err error) {
	return repository.GetMany(params, page, quantity)
}
