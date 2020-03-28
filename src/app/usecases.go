package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

type customerUseCase struct {
	repository CustomerDataHandler
}

func (us *customerUseCase) getValidCustomer(id string) (customer model.Customer, err error) {
	if customer, err = us.repository.Get(id); err == nil {
		err = validateFoundCustomer(customer)
	}
	return customer, err
}

func (us *customerUseCase) Create(customer *model.Customer) error {
	return us.repository.Save(customer)
}

func (us *customerUseCase) Find(id string) (customer model.Customer, err error) {
	return us.getValidCustomer(id)
}

func (us *customerUseCase) Update(id string, customer model.Customer) (newCustomer model.Customer, err error) {
	var foundCustomer model.Customer
	if foundCustomer, err = us.getValidCustomer(id); err != nil {
		return newCustomer, err
	}
	newCustomer = mapCustomerToUpdate(foundCustomer, customer)
	if err = us.repository.Update(newCustomer); err != nil {
		newCustomer = model.Customer{}
	}
	return newCustomer, err
}

func (us *customerUseCase) List(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error) {
	return us.repository.GetMany(params, page, quantity)
}

func (us *customerUseCase) UpdateFromUser(user model.User) (customer model.Customer, err error) {
	if err = validateUser(user); err == nil {
		var val []interface{}
		val = append(val, user.EMail)
		customers, total, err := us.repository.GetMany([]model.SearchParameter{{
			Field:  "eMail", // TODO: realize how to get json tag name from its definition
			Values: val,
		}},
			0,
			2)

		if err == nil {
			if total > 1 {
				err = errors.New("to many register with the same e-mail")
			} else if total == 0 {
				err = errors.New("customer not found")
			} else {
				customer := mapUserToCustomer(user, customers[0])
				err = us.repository.Update(customer)
			}
		}
	}
	return customer, err
}

func NewCustomerUseCase(repository CustomerDataHandler) CustomerUseCase {
	return &customerUseCase{repository: repository}
}
