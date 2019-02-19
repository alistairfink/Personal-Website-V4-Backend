package Orm

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

func NewClient(connectionString string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), connectionString)
	if (err != nil) {
		log.Fatal(err)
	}

	err2 := client.Ping(context.TODO(), readpref.Primary())
	if (err2 != nil) {
		log.Fatal(err)
	}

	return client
}