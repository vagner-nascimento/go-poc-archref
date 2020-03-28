package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type customerRepository struct {
	db  data.NoSqlHandler
	pub amqpPublishHandler
}

func (o *customerRepository) Save(customer *model.Customer) (err error) {
	customer.Id = uuid.New().String()
	if _, err = o.db.InsertOne(customer); err == nil {
		go o.pub.publish(customer)
	}
	return err
}

func (o *customerRepository) Get(id string) (customer model.Customer, err error) {
	if data, err := o.db.FindOne(id); err == nil && data != nil {
		err = unmarshalBsonData(data, &customer)
	}
	return customer, err
}

func (o *customerRepository) GetMany(params []model.SearchParameter, page int64, quantity int64) (customers []model.Customer, total int64, err error) {
	bsonFilters := getBsonFilters(params)
	results := make(chan interface{}, 50)
	if page > 0 {
		page = page - 1
	}

	go o.db.Find(bsonFilters, quantity, page, results, &total)

	for result := range results {
		switch val := result.(type) {
		case []byte:
			var customer model.Customer
			if err = unmarshalBsonData(val, &customer); err != nil {
				return nil, 0, err
			}
			customers = append(customers, customer)
		case error:
			return nil, 0, err
		}
	}
	return customers, total, err
}

func (o *customerRepository) Update(customer model.Customer) error {
	replaceCount, err := o.db.ReplaceOne(bson.M{"id": customer.Id}, customer)
	if err != nil {
		return err
	}
	if replaceCount < 1 {
		return errors.New("none customer was replaced")
	}

	go o.pub.publish(customer)
	return nil
}

func NewCustomerRepository() (custDataHandler app.CustomerDataHandler, err error) {
	var (
		db  data.NoSqlHandler
		pub amqpPublishHandler
	)

	if db, err = data.NewNoSqlDb(config.Get().Data.NoSql.Collections.Customer); err == nil {
		if pub, err = newCustomerPublisher(); err == nil {
			custDataHandler = &customerRepository{
				db:  db,
				pub: pub,
			}
		}
	}
	return custDataHandler, err
}
