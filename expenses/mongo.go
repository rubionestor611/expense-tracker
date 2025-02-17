package expenses

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
}

var (
	instance *Mongo
	once     sync.Once
)

func InitMongo(uri string) *Mongo {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		if err := client.Ping(context.TODO(), nil); err != nil {
			log.Fatalf("Could not ping MongoDB: %v", err)
		}

		instance = &Mongo{client: client}
	})

	return instance
}

func GetMongoInstance() *Mongo {
	if instance == nil {
		log.Fatal("MongoDB is not initialized. Call InitMongo first")
	}

	return instance
}

func GetMongoClient() *mongo.Client {
	return GetMongoInstance().client
}

func ExpensesCollection() *mongo.Collection {
	return GetMongoClient().Database("expenses").Collection("expenses")
}
