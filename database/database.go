package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)
//database global
var (
	Messages *mongo.Collection
	Users    *mongo.Collection
	Channels *mongo.Collection
	Subscription *mongo.Collection
)



func SetupDB()  {

	//db config vars
	var dbHost string = os.Getenv("DB_MONGO_HOST")
	var dbName string = os.Getenv("DB_MONGO_DATABASE")
	clientOptions := options.Client().ApplyURI(dbHost)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	database := client.Database(dbName)
	Messages = database.Collection("messages")
	Users = database.Collection("users")
	Channels = database.Collection("channels")
	Subscription = database.Collection("subscriptions")

}