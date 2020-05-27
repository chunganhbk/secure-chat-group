package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
)
type MongoDataStore struct {
	DBMaster      *mongo.Database
	Session *mongo.Client
	logger  *logrus.Logger
}
const CONNECTED = "Successfully connected to database: %v"

func NewDataStore(logger *logrus.Logger) *MongoDataStore {

	var mongoDataStore *MongoDataStore
	db, session := connect( logger)
	if db != nil && session != nil {

		// log statements here as well
		mongoDataStore = new(MongoDataStore)
		mongoDataStore.DBMaster = db
		mongoDataStore.logger = logger
		mongoDataStore.Session = session
		return mongoDataStore
	}

	logger.Fatalf("Failed to connect to database master")

	return nil
}

func connect(logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(logger)
	})

	return db, session
}

func connectToMongo(logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {

	var err error
	var dbHost string = os.Getenv("DB_MONGO_HOST")
	var dbMaster string = os.Getenv("DB_MONGO_MASTER")
	clientOptions := options.Client().ApplyURI(dbHost)
	session, err := mongo.NewClient(clientOptions)
	if err != nil {
		logger.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	var DB = session.Database(dbMaster)
	logger.Info(CONNECTED, dbMaster)

	return DB, session
}
