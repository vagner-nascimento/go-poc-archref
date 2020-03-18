package app

import "github.com/vagner-nascimento/go-poc-archref/src/model"

type CustomerDataHandler interface {
	Save(customer *model.Customer) error
	Replace(customer model.Customer) error
	Get(id string) (model.Customer, error)
	GetMany(params []SearchParameter, page int64, quantity int64) (customers []model.Customer, total int64, err error)
}
