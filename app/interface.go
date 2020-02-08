package app

type CustomerDataIterable interface {
	Save(c *Customer) error
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
	Update(c *Customer) error
}
