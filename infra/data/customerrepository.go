package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type customerDb struct {
}

func (o *customerDb) Save(c *app.Customer) error {
	client, err := mongoDbClient()
	if err != nil {
		infra.LogError("error on try connect into mongo db", err)
		return connectionError("database server")
	}

	c.Id = uuid.New().String()
	collection := client.Database("golang").Collection("customers")
	_, err = collection.InsertOne(context.TODO(), c)
	if err != nil {
		infra.LogError("error on try insert data on customer collection", err)
		return execError("save", "customer")
	}

	return nil
}

func (o *customerDb) Get(id string) (app.Customer, error) {
	return app.Customer{}, notImplementedError()
}

func (o *customerDb) GetMany(params ...interface{}) ([]app.Customer, error) {
	return []app.Customer{}, notImplementedError()
}

func (o *customerDb) Update(c *app.Customer) error {
	return notImplementedError()
}
