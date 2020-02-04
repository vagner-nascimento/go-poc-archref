package app

type DataIterable interface {
	Save(customer *Customer) (Customer, error)
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
	Update(customer *Customer) (Customer, error)
	Subscribe() (<-chan bool, error)
}
