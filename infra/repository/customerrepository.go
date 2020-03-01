package repository

import (
	"github.com/google/uuid"
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
	"go.mongodb.org/mongo-driver/bson"
)

type CustomerRepository struct {
}

const customerCollection = "customers"

func (o *CustomerRepository) Save(c *app.Customer) error {
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return err
	}

	// TODO: realise how to insert with JSON names declared on app entity
	c.Id = uuid.New().String()
	if _, err = db.Insert(c); err != nil {
		return err
	}

	go publishCustomer(*c)
	return nil
}

func (o *CustomerRepository) Get(id string) (app.Customer, error) {
	var customer app.Customer
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return customer, err
	}

	result, err := db.FindOne(bson.D{{"id", id}})
	if err != nil {
		return customer, err
	}

	customer, err = unmarshalCustomer(result)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (o *CustomerRepository) GetMany(params []app.SearchParameter) ([]app.Customer, error) {
	var filters bson.D

	// TODO: handle operator
	switch len(params) {
	case 0:
		filters = bson.D{{}}
	case 1:
		filters = bson.D{{params[0].Field, params[0].Value}}
	default:
		var ds []bson.D
		for _, param := range params {
			ds = append(ds, bson.D{{param.Field, param.Value}})
		}
		filters = bson.D{{"$and", bson.A{ds}}}

	}

	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return nil, err
	}

	results := make(chan interface{})
	go db.Find(filters, 100, results)

	var customers []app.Customer
	for result := range results {
		switch val := result.(type) {
		case []byte:
			customer, err := unmarshalCustomer(val)
			if err != nil {
				return nil, err
			}
			customers = append(customers, customer)
		case error:
			return nil, val // TODO: handle errors
		}
	}
	return customers, nil
}

// TODO: implement Update
func (o *CustomerRepository) Update(c *app.Customer) error {
	return notImplementedError("customer repository")
}

func unmarshalCustomer(data []byte) (app.Customer, error) {
	var customer app.Customer
	err := bson.Unmarshal(data, &customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
