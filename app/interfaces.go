package app

type CustomerDataHandler interface {
	Save(customer *Customer) error
	Replace(customer Customer) error
	Get(id string) (Customer, error)
	GetMany(params []SearchParameter) ([]Customer, error)
}
