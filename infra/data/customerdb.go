package data

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

// TODO: Implement customerDb to call mongo connection
type customerDb struct {
	db mongo
}

func (o *customerDb) Save(c *app.Customer) error {
	c.Id = "fake_" + uuid.New().String()
	fmt.Println("(fake customerDd) customer save")
	return nil
}

func (o *customerDb) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError()
}

func (o *customerDb) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, notImplementedError()
}

func (o *customerDb) Update(c *app.Customer) error {
	return notImplementedError()
}
