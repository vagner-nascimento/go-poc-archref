package app

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

type CustomerUseCase interface {
	Create(customer *model.Customer) error
	Find(id string) (model.Customer, error)
	Update(id string, customer model.Customer) (model.Customer, error)
	UpdateFromUser(user model.User) (model.Customer, error)
	List(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error)
	UpdateAddress(id string, address model.Address) (model.Customer, error)
}

type customerUseCase struct {
	repository CustomerDataHandler
}

func (uc *customerUseCase) getValidCustomer(id string) (customer model.Customer, err error) {
	if customer, err = uc.repository.Get(id); err == nil {
		err = validateFoundCustomer(customer)
	}
	return customer, err
}

func (uc *customerUseCase) Create(customer *model.Customer) error {
	return uc.repository.Save(customer)
}

func (uc *customerUseCase) Find(id string) (customer model.Customer, err error) {
	return uc.getValidCustomer(id)
}

func (uc *customerUseCase) Update(id string, customer model.Customer) (newCustomer model.Customer, err error) {
	var foundCustomer model.Customer
	if foundCustomer, err = uc.getValidCustomer(id); err != nil {
		return newCustomer, err
	}
	newCustomer = mapCustomerToUpdate(foundCustomer, customer)
	if err = uc.repository.Update(newCustomer); err != nil {
		newCustomer = model.Customer{}
	}
	return newCustomer, err
}

func (uc *customerUseCase) List(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error) {
	return uc.repository.GetMany(params, page, quantity)
}

func (uc *customerUseCase) UpdateFromUser(user model.User) (customer model.Customer, err error) {
	if err = validateUser(user); err == nil {
		var val []interface{}
		val = append(val, user.EMail)
		customers, total, err := uc.repository.GetMany([]model.SearchParameter{{
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
				err = uc.repository.Update(customer)
			}
		}
	}
	return customer, err
}

func (uc *customerUseCase) UpdateAddress(id string, address model.Address) (customer model.Customer, err error) {
	if customer, err = uc.getValidCustomer(id); err == nil {
		setCustomerAddress(&customer, address)
		err = uc.repository.Update(customer)
	}
	return customer, err
}

func NewCustomerUseCase(repository CustomerDataHandler) CustomerUseCase {
	return &customerUseCase{repository: repository}
}
