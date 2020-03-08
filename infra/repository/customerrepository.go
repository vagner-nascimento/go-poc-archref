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

func (o *CustomerRepository) GetMany(params []app.SearchParameter, page int64, quantity int64) (customers []app.Customer, total int64, err error) {
	var filters bson.D
	filters, err = getBsonFilters(params)
	if err != nil {
		return
	}

	var db *data.MongoDb
	db, err = data.NewMongoDb(customerCollection)
	if err != nil {
		return
	}

	results := make(chan interface{}, 50)
	go db.Find(filters, quantity, page, results, &total)

	for result := range results {
		switch val := result.(type) {
		case []byte:
			var customer app.Customer
			customer, err = unmarshalCustomer(val)
			if err != nil {
				return
			}
			customers = append(customers, customer)
		case error:
			return // TODO: handle errors
		}
	}

	return customers, total, err
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
		return errors.New("none Customer was replaced")
	}

	go publishCustomer(customer)

	return nil
}

func getBsonFilters(params []app.SearchParameter) (filters bson.D, err error) {
paramsSwitch:
	switch len(params) {
	case 0:
		filters = bson.D{{}}
	case 1: // TODO: when came from QUERY PARAM is lower case. Realise how to find it into Customer that is like this: Name, Id
		field, ok := reflect.TypeOf(&app.Customer{}).Elem().FieldByName(params[0].Field)
		if !ok {
			err = errors.New(fmt.Sprintf("field %s not found on Customer", params[0].Field))
			break
		}

		key := field.Tag.Get("bson")
		if len(key) <= 0 {
			err = errors.New(fmt.Sprintf("field %s has no bson tag into customer", params[0].Field))
			break
		}

		filters = bson.D{{key, params[0].Value}}
	default:
		var andFilter []bson.D
		for _, param := range params {
			field, ok := reflect.TypeOf(app.Customer{}).Elem().FieldByName(param.Field)
			if !ok {
				err = errors.New(fmt.Sprintf("field %s not found on Customer", param.Field))
				break paramsSwitch
			}

			key := field.Tag.Get("bson")
			if len(key) <= 0 {
				err = errors.New(fmt.Sprintf("field %s has not bson tag into customer", param.Field))
				break paramsSwitch
			}

			andFilter = append(andFilter, bson.D{{key, param.Value}})
		}

		filters = bson.D{{"$and", bson.A{andFilter}}}
	}

	return filters, err
}

func unmarshalCustomer(data []byte) (app.Customer, error) {
	var customer app.Customer
	err := bson.Unmarshal(data, &customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
