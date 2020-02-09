package data

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

type customerDb struct {
}

// TODO: properly implements mongo connection
// TODO: realise why name, id and other properties was not save into database
func (o *customerDb) Save(c *app.Customer) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		fmt.Println("Error on mongo connections")
		fmt.Println(err)
		return err
	}
	defer client.Disconnect(ctx)

	c.Id = uuid.New().String()
	collection := client.Database("golang").Collection("customers")
	res, err := collection.InsertOne(ctx, c)
	if err != nil {
		fmt.Println("Error on insert on collection")
		fmt.Println(err)
		return err
	}

	fmt.Println("Data inserted. insert response")
	fmt.Println(res)
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
