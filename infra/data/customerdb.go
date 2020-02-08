package data

import (
	"fmt"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

// TODO: Implement customerDb to call mongo connection
type customerDb struct {
	db mongo
}

func (m *customerDb) Save(c *app.Customer) error {
	c.Id = "newUuid"
	fmt.Println("(fake customerDd) customer save")
	return nil
}

func (m *customerDb) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError()
}

func (m *customerDb) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, notImplementedError()
}

func (m *customerDb) Update(c *app.Customer) error {
	return notImplementedError()
}
