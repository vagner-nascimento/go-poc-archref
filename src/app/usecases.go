package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

type customerUseCase struct {
	repository CustomerDataHandler
}

func (us *customerUseCase) Create(customer *model.Customer) error {
	return us.repository.Save(customer)
}

func (us *customerUseCase) Find(id string) (customer model.Customer, err error) {
	if customer, err = us.repository.Get(id); err == nil {
		if len(customer.Id) <= 0 {
			customer = model.Customer{}
			err = errors.New("customer not found")
		}
	}
	return customer, err
}

func (us *customerUseCase) Update(id string, customer model.Customer) (newCustomer model.Customer, err error) {
	var foundCustomer model.Customer
	foundCustomer, err = us.repository.Get(id)
	if err != nil {
		return newCustomer, err
	}

	newCustomer = mapCustomerToUpdate(foundCustomer, customer)
	if err == nil {
		err = us.repository.Update(newCustomer)
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
