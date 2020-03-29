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
	db  data.NoSqlHandler
	pub amqpPublishHandler
}

func (repo *supplierRepository) Save(sup *model.Supplier) (err error) {
	sup.Id = uuid.New().String()
	if _, err = repo.db.InsertOne(sup); err == nil {
		go repo.pub.publish(sup)
	}
	return err
}

func (repo *supplierRepository) Update(sup model.Supplier) (err error) {
	if replaceCount, err := repo.db.ReplaceOne(bson.M{"id": sup.Id}, sup); err == nil && replaceCount < 1 {
		err = errors.New("none supplier was replaced")
	}
	if err == nil {
		go repo.pub.publish(sup)
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
	var (
		db  data.NoSqlHandler
		pub amqpPublishHandler
	)

	if db, err = data.NewNoSqlDb(config.Get().Data.NoSql.Collections.Supplier); err == nil {
		if pub, err = newSupplierPublisher(); err == nil {
			supDataHandler = &supplierRepository{
				db:  db,
				pub: pub,
			}
		}
	}
	return supDataHandler, err
}
