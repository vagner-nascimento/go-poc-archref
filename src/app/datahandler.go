package app

import "github.com/vagner-nascimento/go-poc-archref/src/model"

type CustomerDataHandler interface {
	Save(customer *model.Customer) error
	Update(customer model.Customer) error
	Get(id string) (model.Customer, error)
	GetMany(params []model.SearchParameter, page int64, quantity int64) ([]model.Customer, int64, error)
}

type SupplierDataHandler interface {
	Save(sup *model.Supplier) error
	Update(sup model.Supplier) error
	Get(id string) (model.Supplier, error)
	GetMany(params []model.SearchParameter, page int64, quantity int64) ([]model.Supplier, int64, error)
}
