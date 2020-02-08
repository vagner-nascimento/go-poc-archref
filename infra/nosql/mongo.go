package nosql

import (
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/app"
)

type mongo struct {
}

func (m *mongo) Save(c *app.Customer) error {
	c.Id = "newUuid"
	fmt.Println("Saved on mongo")
	return nil
}

func (m *mongo) Get(id string) (app.Customer, error) {
	c, _ := app.NewCustomer(m) // use NEW dps

	c.Alias = "alias"
	c.CreditCardHash = "hashCard"
	c.Name = "fake data"
	c.Id = id

	return c, nil
}

func (m *mongo) GetMany(params ...interface{}) ([]app.Customer, error) {
	var cs []app.Customer

	c, _ := app.NewCustomer(m)
	c.Alias = "alias"
	c.CreditCardHash = "hashCard"
	c.Name = "fake data"
	c.Id = "id"

	c1, _ := app.NewCustomer(m)
	c1.Alias = "alias1"
	c1.CreditCardHash = "hashCard1"
	c1.Name = "fake data1"
	c1.Id = "id1"

	cs = append(cs, c)
	cs = append(cs, c1)

	return cs, nil
}

func (m *mongo) Update(c *app.Customer) error {
	fmt.Println("Updated on mongo")
	return nil
}

func New() *mongo {
	return &mongo{}
}
