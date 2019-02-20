package Orm

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	
	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func AddEducation(db *mongo.Database, config *Config.Config) *Models.Education {
	return nil
}

func GetEducation(db *mongo.Database, config *Config.Config) *Models.Education {
	return nil
}

func GetAllEducation(db *mongo.Database, config *Config.Config) *[]Models.Education {
	collection := db.Collection("Education")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Fatal(err)
	}

	var educations []Models.Education
	for (curr.Next(context.TODO())) {
		elem := Models.Education{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Fatal(errr)
		}

		educations = append(educations, elem)
	}

	return &educations
}

func EditEducation(db *mongo.Database, config *Config.Config) *Models.Education {
	return nil
}

func DeleteEducation(db *mongo.Database, config *Config.Config) bool {
	return true
}