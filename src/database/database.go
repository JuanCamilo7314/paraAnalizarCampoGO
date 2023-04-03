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

func InitMongoConnection() error {
	urlDb := os.Getenv("DATABASE_URL")
	nameDb := os.Getenv("DATABASE_NAME")
	var err error

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

	return err
}

func GetCollection(collection string) *mongo.Collection {
	return Db.database.Collection(collection)
}

func (m *MongoDB) CloseConnection() error {
	return m.client.Disconnect(context.Background())
}
