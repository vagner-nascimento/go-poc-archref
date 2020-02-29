package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
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

	if _, err = db.Insert(c); err != nil {
		return err
	}

	go publishCustomer(*c)
	return nil
}

func (o *CustomerRepository) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError("customer repository")
}

// TODO: implement GetMany
func (o *CustomerRepository) GetMany(params []app.SearchParameter) ([]app.Customer, error) {
	var filters bson.D

	// TODO: handle operator
	switch len(params) {
	case 0:
		filters = bson.D{{}}
	case 1:
		filters = bson.D{{params[0].Field, params[0].Value}}
	default:
		{
			var ds []bson.D
			for _, param := range params {
				ds = append(ds, bson.D{{param.Field, param.Value}})
			}
			filters = bson.D{{"$and", bson.A{ds}}}
		}
	}

	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return nil, err
	}

	res, err := db.Find(filters, 100, reflect.TypeOf(app.Customer{}))
	if err != nil {
		return nil, err
	}

	var customers []app.Customer
	for _, item := range res {
		c, ok := item.(app.Customer) // TODO: Realise how to convert to Customer
		if !ok {
			fmt.Println("error on convert customer")
			return nil, errors.New("error on convert customer")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

// TODO: implement Update
func (o *CustomerRepository) Update(c *app.Customer) error {
	return notImplementedError("customer repository")
}
