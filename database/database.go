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
var Site *mongo.Collection

func SetupDB() {

	//db config vars
	var dbHost string = os.Getenv("DB_MONGO_HOST")
	var dbMaster string = os.Getenv("DB_MONGO_MASTER")
	clientOptions := options.Client().ApplyURI(dbHost)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	database := client.Database(dbMaster)
	Site = database.Collection("sites")
	//Users = database.Collection("users")
	//Channels = database.Collection("channels")
	//Subscription = database.Collection("subscriptions")
}
