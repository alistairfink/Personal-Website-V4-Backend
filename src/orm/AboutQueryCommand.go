package Orm

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func GetAbout(db *mongo.Client, config *Config.Config) *Models.About {
	collection := db.Database(config.Mongo.Name).Collection("About")
	var result Models.About
	//filter := Models.About{}
	// TODO: Figure out filter
	err := collection.FindOne(context.TODO(), nil).Decode(&result)
	if (err != nil) {
		log.Fatal(err)
	}

	return &result 
}

func EditAbout() *Models.About {
	return nil
}