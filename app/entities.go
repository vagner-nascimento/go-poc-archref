package app

type Customer struct {
	Id             string `id: "id"`
	Name           string `name: "name"`
	EMail          string `eMail: "eMail"`
	CreditCardHash string
	BirthYear      int `birthYear: "birthYear"`
	BirthDay       int `birthDay: "birthDay"`
	BirthMont      int `birthMont: "birthMonth"`
	data           CustomerDataHandler
}

func (c *Customer) save() error {
	return c.data.Save(c)
}

type User struct {
	Id       string `id: "id"`
	Name     string `name: "name"`
	EMail    string `eMail: "eMail"`
	Customer Customer
	UseName  string `userName: "userName"`
	Password string
	data     UserDataHandler
}

func NewCustomer(db CustomerDataHandler) *Customer {
	return &Customer{data: db}
}

func NewUser(db UserDataHandler) *User {
	return &User{data: db}
}
