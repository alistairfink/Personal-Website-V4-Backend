package Orm

import (
	"context"
	"log"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func AddExperience(db *mongo.Database, config *Config.Config) *Models.Experience {
	return nil
}

func GetExperience(db *mongo.Database, config *Config.Config) *Models.Experience {
	return nil
}

func GetAllExperience(db *mongo.Database, config *Config.Config) *[]Models.Experience {
	collection := db.Collection("Experience")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Fatal(err)
	}

	var experiences []Models.Experience
	for (curr.Next(context.TODO())) {
		elem := Models.Experience{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Fatal(errr)
		}

		experiences = append(experiences, elem)
	}

	return &experiences
}

func EditExperience(db *mongo.Database, config *Config.Config) *Models.Experience {
	return nil
}

func DeleteExperience(db *mongo.Database, config *Config.Config) bool {
	return true
}