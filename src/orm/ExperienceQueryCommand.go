package Orm

import (
	"context"
	"log"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func AddExperience(db *mongo.Database, config *Config.Config, newExperience *Models.CreateExperience) *Models.Experience {
	collection := db.Collection("Experience")
	res, err := collection.InsertOne(context.TODO(), newExperience)
	if (err != nil) {
		log.Println(err)
		return nil
	}

	result := GetExperience(db, config, res.InsertedID.(primitive.ObjectID).Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func GetExperience(db *mongo.Database, config *Config.Config, id string) *Models.Experience {
	collection := db.Collection("Experience")
	var result Models.Experience
	_id, iderr := primitive.ObjectIDFromHex(id)
	if (iderr != nil) {
		log.Println(iderr)
		return nil
	}

	err := collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&result)
	if (err != nil) {
		log.Println(err)
		return nil
	}

	return &result
}

func GetAllExperience(db *mongo.Database, config *Config.Config) *[]Models.Experience {
	collection := db.Collection("Experience")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Println(err)
		return nil
	}

	var experiences []Models.Experience
	for (curr.Next(context.TODO())) {
		elem := Models.Experience{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Println(errr)
			return nil
		}

		experiences = append(experiences, elem)
	}

	return &experiences
}

func EditExperience(db *mongo.Database, config *Config.Config, editExperience *Models.Experience) *Models.Experience {
	collection := db.Collection("Experience")
	exists := GetExperience(db, config, editExperience.Id.Hex())
	if (exists == nil) {
		return nil
	}

	res := collection.FindOneAndReplace(context.TODO(), bson.M{"_id": editExperience.Id}, editExperience)
	if (res.Err() != nil) {
		log.Println(res.Err())
		return nil
	}

	result := GetExperience(db, config, editExperience.Id.Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func DeleteExperience(db *mongo.Database, config *Config.Config, id string) bool {
	collection := db.Collection("Experience")
	_id, iderr := primitive.ObjectIDFromHex(id)
	if (iderr != nil) {
		log.Println(iderr)
		return false
	}

	res := GetExperience(db, config, id)
	if (res == nil) {
		return false
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if (err != nil) {
		log.Println(err)
		return false
	}

	return true
}
