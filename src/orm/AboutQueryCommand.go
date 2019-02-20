package Orm

import (
	"context"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func GetAbout(db *mongo.Client, config *Config.Config) *Models.About {
	collection := db.Database(config.Mongo.Name).Collection("About")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Fatal(err)
	}

	var abouts []Models.About
	for (curr.Next(context.TODO())) {
		elem := Models.About{}
		err := curr.Decode(&elem)
		if (err != nil) {
			log.Fatal(err)
		}

		abouts = append(abouts, elem)
	}

	return &abouts[0]
}

func EditAbout() *Models.About {
	return nil
}