package repository

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type CustomerRepository struct {
}

const customerCollection = "customers"

func (o *CustomerRepository) Save(customer *app.Customer) error {
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return err
	}

	customer.Id = uuid.New().String()
	if _, err = db.Insert(customer); err != nil {
		return err
	}

	go publishCustomer(*customer)
	return nil
}

func (o *CustomerRepository) Get(id string) (app.Customer, error) {
	var customer app.Customer
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return customer, err
	}

	result, err := db.FindOne(bson.D{{"id", id}})
	if err != nil || result == nil {
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
		field, ok := reflect.TypeOf(&app.Customer{}).Elem().FieldByName(params[0].Field)
		if !ok {
			return nil, errors.New(fmt.Sprintf("field %s not found on Customer", params[0].Field))
		}

		key := field.Tag.Get("bson")
		if len(key) <= 0 {
			return nil, errors.New(fmt.Sprintf("field %s has not bson tag into customer", params[0].Field))
		}

		filters = bson.D{{key, params[0].Value}}
	default:
		var ds []bson.D
		for _, param := range params {
			field, ok := reflect.TypeOf(app.Customer{}).Elem().FieldByName(param.Field)
			if !ok {
				return nil, errors.New(fmt.Sprintf("field %s not found on Customer", param.Field))
			}

			key := field.Tag.Get("bson")
			if len(key) <= 0 {
				return nil, errors.New(fmt.Sprintf("field %s has not bson tag into customer", param.Field))
			}

			ds = append(ds, bson.D{{key, param.Value}})
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

func (o *CustomerRepository) UpdateMany(param app.SearchParameter, data []app.UpdateParameter) (int64, error) {
	return 0, notImplementedError("customer repository")
}

func (o *CustomerRepository) Replace(customer app.Customer) error {
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return err
	}

	replaceCount, err := db.ReplaceOne(bson.M{"id": customer.Id}, customer)
	if err != nil {
		return err
	}

	if replaceCount < 1 {
		return errors.New("none Customer was replace")
	}

	go publishCustomer(customer)

	return nil
}

func (o *CustomerRepository) Update(id string, data []app.UpdateParameter) error {

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
