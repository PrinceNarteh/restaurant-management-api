package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection *mongo.Collection
	FoodCollection *mongo.Collection
	MenuCollection *mongo.Collection
)

func Connect() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// defer disconnect
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Database connected successfully!")

	DB := client.Database("restaurant-management")
	UserCollection = DB.Collection("users")
	FoodCollection = DB.Collection("foods")
	MenuCollection = DB.Collection("menus")
}
