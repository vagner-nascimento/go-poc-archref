package app

type person struct {
	Id   string `id: "id"`
	Name string `name: "name"`
}

type Customer struct {
	person
	Alias          string `alias: "alias"`
	CreditCardHash string
}

type User struct {
	person
	UseName  string `userName: "userName"`
	Password string
}
