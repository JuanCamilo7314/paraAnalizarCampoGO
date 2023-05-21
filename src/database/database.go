package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Db   *MongoDB
	once sync.Once
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

/*
	InitMongoConnection initializes the connection to the MongoDB database.

It utilizes the singleton pattern to ensure that only one connection is created.
Once the connection is established, the function assigns the client and
database objects to the global Db variable.
*/
func InitMongoConnection() {
	urlDb := os.Getenv("DATABASE_URL")
	nameDb := os.Getenv("DATABASE_NAME")

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(urlDb)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		database := client.Database(nameDb)
		Db = &MongoDB{
			client:   client,
			database: database,
		}
	})
}

func (db *MongoDB) GetCollection(collection string) *mongo.Collection {
	return db.database.Collection(collection)
}

func (db *MongoDB) CloseConnection() error {
	return db.client.Disconnect(context.Background())
}
