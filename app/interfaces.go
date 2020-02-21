package app

type CustomerDataHandler interface {
	Save(c *Customer) error
	Update(c *Customer) error
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
}
