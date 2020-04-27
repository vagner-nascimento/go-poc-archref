package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

type supplierRepository struct {
	db       data.NoSqlHandler
	pubTopic string
}

func (repo *supplierRepository) Save(sup *model.Supplier) (err error) {
	sup.Id = uuid.New().String()
	if _, err = repo.db.InsertOne(sup); err == nil {
		go publishMessage(sup, repo.pubTopic)
	}

	return err
}

func (repo *supplierRepository) Update(sup model.Supplier) (err error) {
	// TODO: review it, is strange
	if replaceCount, err := repo.db.ReplaceOne(bson.M{"id": sup.Id}, sup); err == nil && replaceCount < 1 {
		err = errors.New("none supplier was replaced")
	} else {
		go publishMessage(&sup, repo.pubTopic)
	}

	return err
}

func (repo *supplierRepository) Get(id string) (sup model.Supplier, err error) {
	if data, err := repo.db.FindOne(id); err == nil && data != nil {
		err = unmarshalBsonData(data, &sup)
	}
	return sup, err
}

func (repo *supplierRepository) GetMany(params []model.SearchParameter, page int64, quantity int64) (sups []model.Supplier, total int64, err error) {
	bsonFilters := getBsonFilters(params)
	results := make(chan interface{}, 50)
	if page > 0 {
		page = page - 1
	}

	go repo.db.Find(bsonFilters, quantity, page, results, &total)
	// TODO Think how to make this result pik generic
	for result := range results {
		switch val := result.(type) {
		case []byte:
			var sup model.Supplier
			if err = unmarshalBsonData(val, &sup); err != nil {
				return nil, 0, err
			}
			sups = append(sups, sup)
		case error:
			return nil, 0, err
		}
	}

	return sups, total, err
}

func NewSupplierRepository() (supDataHandler app.SupplierDataHandler, err error) {
	var db data.NoSqlHandler
	if db, err = data.NewNoSqlDb(config.Get().Data.NoSql.Collections.Supplier); err == nil {
		supDataHandler = &supplierRepository{
			db:       db,
			pubTopic: config.Get().Integration.Amqp.Pubs.Supplier.Topic,
		}
	}
	return supDataHandler, err
}
