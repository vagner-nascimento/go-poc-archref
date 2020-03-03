package app

import (
	"reflect"
)

func CreateCustomer(customerData []byte, repository CustomerDataHandler) (interface{}, error) {
	c, err := getCustomer(customerData)
	if err == nil {
		if err = setCustomerCreditCardHash(&c); err == nil {
			err = repository.Save(&c)
		}
	}

	return c, err
}

func UpdateCustomerFromUser(userData []byte, repository CustomerDataHandler) (Customer, error) {
	user, err := getUser(userData)
	if err != nil {
		return Customer{}, err
	}

	if err = validateUser(user); err != nil {
		return Customer{}, err
	}

	customers, err := repository.GetMany([]SearchParameter{{
		Field:    "EMail",
		Operator: "=",
		Value:    user.EMail,
	}})
	if err != nil {
		return Customer{}, err
	}

	var foundCustomer Customer
	if len(customers) > 0 {
		foundCustomer = customers[0]
	}

	if foundCustomer.Id == "" {
		return Customer{}, notFoundError(reflect.TypeOf(Customer{}))
	}

	newCustomer := mergeUserToCustomer(user, foundCustomer)
	if err = repository.Replace(newCustomer); err != nil {
		return Customer{}, err
	}

	return newCustomer, nil
}

func FindCustomer(id string, repository CustomerDataHandler) (Customer, error) {
	if customer, err := repository.Get(id); err != nil {
		return Customer{}, err
	} else if len(customer.Id) <= 0 {
		return Customer{}, notFoundError(reflect.TypeOf(Customer{}))
	} else {
		return customer, nil
	}
}
