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

func AddEducation(db *mongo.Database, config *Config.Config, newEducation *Models.CreateEducation) *Models.Education {
	collection := db.Collection("Education")
	res, err := collection.InsertOne(context.TODO(), newEducation)
	if (err != nil) {
		log.Println(err)
		return nil
	}

	result := GetEducation(db, config, res.InsertedID.(primitive.ObjectID).Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func GetEducation(db *mongo.Database, config *Config.Config, id string) *Models.Education {
	collection := db.Collection("Education")
	var result Models.Education
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

func GetAllEducation(db *mongo.Database, config *Config.Config) *[]Models.Education {
	collection := db.Collection("Education")
	curr, err := collection.Find(context.TODO(), bson.M{})
	if (err != nil) {
		log.Println(err)
		return nil
	}

	var educations []Models.Education
	for (curr.Next(context.TODO())) {
		elem := Models.Education{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Println(errr)
			return nil
		}

		educations = append(educations, elem)
	}

	return &educations
}

func EditEducation(db *mongo.Database, config *Config.Config, editEducation *Models.Education) *Models.Education {
	collection := db.Collection("Education")
	exists := GetEducation(db, config, editEducation.Id.Hex())
	if (exists == nil) {
		return nil
	}

	res := collection.FindOneAndReplace(context.TODO(), bson.M{"_id": editEducation.Id}, editEducation)
	if (res.Err() != nil) {
		log.Println(res.Err())
		return nil
	}

	result := GetEducation(db, config, editEducation.Id.Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func DeleteEducation(db *mongo.Database, config *Config.Config, id string) bool {
	collection := db.Collection("Education")
	_id, iderr := primitive.ObjectIDFromHex(id)
	if (iderr != nil) {
		log.Println(iderr)
		return false
	}

	res := GetEducation(db, config, id)
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