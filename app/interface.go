package app

type CustomerDataIterable interface {
	Save(customer *Customer) (Customer, error)
	Get(id string) (Customer, error)
	GetMany(params ...interface{}) ([]Customer, error)
	Update(customer *Customer) (Customer, error)
}

type UserDataIterable interface {
	Save(user *User) (User, error)
	Get(id string) (User, error)
	GetMany(params ...interface{}) ([]User, error)
	Update(user *User) (User, error)
}
