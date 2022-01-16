package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx    context.Context
	DB     *mongo.Database
	Client *mongo.Client
	User   *mongo.Collection
)

func Connect() {
	var cancel context.CancelFunc
	Ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("couldn't connect to database")
	}

	if err = Client.Connect(Ctx); err != nil {
		log.Fatal("couldn't connect to database")
	} else {
		fmt.Println("Database connected successfully!")
	}
	defer Client.Disconnect(Ctx)

	DB = Client.Database("restaurant-management")
	User = DB.Collection("users")
}
