package app

import "github.com/vagner-nascimento/go-poc-archref/src/model"

type CustomerDataHandler interface {
	Save(customer *model.Customer) error
	Update(customer model.Customer) error
	Get(id string) (model.Customer, error)
	GetMany(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error)
}

type CustomerUseCase interface {
	Create(customer *model.Customer) error
	Find(id string) (model.Customer, error)
	Update(id string, customer model.Customer) (model.Customer, error)
	UpdateFromUser(user model.User) (model.Customer, error)
	List(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error)
}
