package Orm

import (
	"context"
	"log"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func GetAbout(db *mongo.Database, config *Config.Config) *Models.About {
	collection := db.Collection("About")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Println(err)
		return nil
	}

	var abouts []Models.About
	for (curr.Next(context.TODO())) {
		elem := Models.About{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Println(errr)
			return nil
		}

		abouts = append(abouts, elem)
	}

	return &abouts[0]
}

func EditAbout(db *mongo.Database, config *Config.Config, aboutEdit *Models.About) *Models.About {
	collection := db.Collection("About")
	res := collection.FindOneAndReplace(context.TODO(), bson.M{"_id": aboutEdit.Id}, aboutEdit)
	if (res.Err() != nil) {
		log.Println(res.Err())
		return nil
	}

	result := GetAbout(db, config)
	if (result == nil) {
		return nil
	}

	return result
}