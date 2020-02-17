package repository

import (
	"github.com/google/uuid"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type customerRepository struct {
}

const customerCollection = "customers"

func (o *customerRepository) Save(c *app.Customer) error {
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return err
	}

	c.Id = uuid.New().String()

	_, err = db.Insert(c)
	if err != nil {
		return err
	}

	return nil
}

// TODO: implementation of Customer's repository
func (o *customerRepository) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError("customer repository")
}

func (o *customerRepository) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, notImplementedError("customer repository")
}

func (o *customerRepository) Update(c *app.Customer) error {
	return notImplementedError("customer repository")
}
