package app

type CustomerDataHandler interface {
	Save(customer *Customer) error
	Update(id string, data []UpdateParameter) error
	UpdateMany(param SearchParameter, data []UpdateParameter) (int64, error)
	Replace(customer Customer) error
	Get(id string) (Customer, error)
	GetMany(params []SearchParameter) ([]Customer, error)
}
