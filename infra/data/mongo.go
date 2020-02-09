package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDb struct {
}

func tst() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Error on mongo connections")
		fmt.Println(err)
		panic(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("golang").Collection("customers")

	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		fmt.Println("Error on insert on collection")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("insert response")
	fmt.Println(res)
}
