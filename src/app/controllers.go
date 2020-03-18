package app

import "errors"

func CreateCustomer(customerData []byte, repository CustomerDataHandler) (Customer, error) {
	customer, err := getCustomer(customerData)
	if err == nil {
		if err = repository.Save(&customer); err != nil {
			customer = Customer{}
		}
	}

	return customer, err
}

func UpdateCustomerEmail(id string, customerData []byte, repository CustomerDataHandler) (newCustomer Customer, err error) {
	foundCustomer, err := repository.Get(id)
	if err != nil {
		return newCustomer, err
	}

	if newCustomer, err = getCustomerToUpdateEmail(foundCustomer, customerData); err != nil {
		return newCustomer, err
	}

	if err = repository.Replace(newCustomer); err != nil {
		newCustomer = Customer{}
		return newCustomer, err
	}

	return newCustomer, err
}

func UpdateCustomer(id string, customerData []byte, repository CustomerDataHandler) (newCustomer Customer, err error) {
	var foundCustomer Customer
	foundCustomer, err = repository.Get(id)
	if err != nil {
		return newCustomer, err
	}

	newCustomer, err = getCustomerToUpdate(foundCustomer, customerData)
	if err != nil {
		return newCustomer, err
	}

	if err = repository.Replace(newCustomer); err != nil {
		newCustomer = Customer{}
		return newCustomer, err
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

	var val []interface{}
	val = append(val, user.EMail)
	customers, total, err := repository.GetMany([]SearchParameter{{
		Field:  "eMail", // TODO: realize how to get json tag name from its definition
		Values: val,
	}},
		0,
		2)

	if total > 1 {
		return Customer{}, errors.New("to many register with the same e-mail")
	} else if total == 0 {
		return Customer{}, errors.New("customer not found")
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
		return Customer{}, errors.New("customer not found")
	} else {
		return customer, nil
	}
}

func FindCustomers(params []SearchParameter, page int64, quantity int64, repository CustomerDataHandler) (res []Customer, total int64, err error) {
	return repository.GetMany(params, page, quantity)
}
