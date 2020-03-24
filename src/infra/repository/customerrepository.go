package repository

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type customerRepository struct {
}

const customerCollection = "customers"

func (o *customerRepository) Save(customer *model.Customer) error {
	db, err := data.NewMongoDb(customerCollection)
	if err != nil {
		return err
	}

	customer.Id = uuid.New().String()
	if _, err = db.Insert(customer); err != nil {
		return err
	}

	go publishCustomer(*customer) // TODO: remove it from here
	return nil
}

func (o *customerRepository) Get(id string) (model.Customer, error) {
	var customer model.Customer
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

func (o *customerRepository) GetMany(params []model.SearchParameter, page int64, quantity int64) (customers []model.Customer, total int64, err error) {
	var db *data.MongoDb
	db, err = data.NewMongoDb(customerCollection)
	if err != nil {
		return nil, 0, err
	}

	bsonFilters := getBsonFilters(params)
	results := make(chan interface{}, 50)
	if page > 0 {
		page = page - 1
	}
	go db.Find(bsonFilters, quantity, page, results, &total)

	for result := range results {
		switch val := result.(type) {
		case []byte:
			var customer model.Customer
			customer, err = unmarshalCustomer(val)
			if err != nil {
				return nil, 0, err
			}
			customers = append(customers, customer)
		case error:
			return nil, 0, err // TODO: handle errors
		}
	}

	return customers, total, err
}

func (o *customerRepository) Update(customer model.Customer) error {
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

func NewCustomerRepository() app.CustomerDataHandler {
	return &customerRepository{}
}

// TODO: think in a better place form these 2 funcs:
func getBsonFilters(params []model.SearchParameter) bson.D {
	convertValue := func(val interface{}) (res interface{}) {
		if res, err := strconv.ParseInt(val.(string), 0, 64); err == nil {
			return res
		}
		if res, err := strconv.ParseFloat(val.(string), 64); err == nil {
			return res
		}
		if res, err := strconv.ParseBool(val.(string)); err == nil {
			return res
		}

		return val.(string)
	}
	convertValues := func(vals []interface{}) (res []interface{}) {
		for _, val := range vals {
			res = append(res, convertValue(val))
		}
		return res
	}

	getBsonD := func(param model.SearchParameter) bson.D {
		bsonD := bson.D{{}}
		if len(param.Values) > 0 {
			if len(param.Values) == 1 {
				bsonD = bson.D{{param.Field, convertValue(param.Values[0])}}
			} else {
				bsonD = bson.D{{param.Field, bson.M{"$in": convertValues(param.Values)}}}
			}
		}
		return bsonD
	}

	var filters bson.D
	switch len(params) {
	case 0:
		filters = bson.D{{}}
	case 1:
		filters = getBsonD(params[0])
	default: //
		var andFilter []bson.D
		for _, param := range params {
			andFilter = append(andFilter, getBsonD(param))
		}
		filters = bson.D{{"$and", andFilter}}
	}

	return filters
}

func unmarshalCustomer(data []byte) (model.Customer, error) {
	var customer model.Customer
	err := bson.Unmarshal(data, &customer)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
