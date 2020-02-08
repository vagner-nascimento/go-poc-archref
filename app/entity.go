package app

type person struct {
	Id   string `id: "id"`
	Name string `name: "name"`
}

type Customer struct {
	person
	CreditCardHash string
	data           CustomerDataHandler
}

func (c *Customer) save() error {
	return c.data.Save(c)
}

type User struct {
	person
	UseName  string `userName: "userName"`
	Password string
}

func NewCustomer(db CustomerDataHandler) *Customer {
	return &Customer{data: db}
}
