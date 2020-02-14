package repository

import (
	"github.com/google/uuid"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type customerRepository struct {
}

// TODO: make a package named repository and put repos into it (repo will depends on data)
// TODO: Realise how to put collections Name into only one place
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
	return app.Customer{}, data.NotImplementedError()
}

func (o *customerRepository) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, data.NotImplementedError()
}

func (o *customerRepository) Update(c *app.Customer) error {
	return data.NotImplementedError()
}
