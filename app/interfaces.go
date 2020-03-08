package app

type CustomerDataHandler interface {
	Save(customer *Customer) error
	Replace(customer Customer) error
	Get(id string) (Customer, error)
	GetMany(params []SearchParameter, page int64, quantity int64) (customers []Customer, total int64, err error)
}
