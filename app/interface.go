package app

type CustomerDataIterable interface {
	Save(c *Customer) error
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
	Update(c *Customer) error
}

type UserDataIterable interface {
	Save(user *User) error
	Get(id string) (User, error)
	GetMany(params ...interface{}) ([]User, error)
	Update(user *User) error
}
