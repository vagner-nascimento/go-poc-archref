package app

type CustomerDataHandler interface {
	Save(c *Customer) error
	Update(c *Customer) error
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
}

type UserDataHandler interface {
	Save(u *User) error
	Update(u *User) error
	Get(id string) (User, error)
	GetMany(params ...interface{}) ([]User, error)
}
