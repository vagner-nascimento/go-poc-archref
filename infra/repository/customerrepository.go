package repository

import (
	"github.com/google/uuid"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type CustomerRepository struct {
}

const customerCollection = "customers"

func (o *CustomerRepository) Save(c *app.Customer) error {
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

// TODO: implements all Customer's repository methods
func (o *CustomerRepository) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError("customer repository")
}

func (o *CustomerRepository) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, notImplementedError("customer repository")
}

func (o *CustomerRepository) Update(c *app.Customer) error {
	return notImplementedError("customer repository")
}
