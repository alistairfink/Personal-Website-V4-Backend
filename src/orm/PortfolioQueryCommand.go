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

func AddPortfolio(db *mongo.Database, config *Config.Config, newModel *Models.CreatePortfolio) *Models.Portfolio {
	collection := db.Collection("Portfolio")
	res, err := collection.InsertOne(context.TODO(), newModel)
	if (err != nil) {
		log.Println(err)
		return nil
	}

	result := GetPortfolio(db, config, res.InsertedID.(primitive.ObjectID).Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func GetPortfolio(db *mongo.Database, config *Config.Config, id string) *Models.Portfolio {
	collection := db.Collection("Portfolio")
	var result Models.Portfolio
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

func GetAllPortfolio(db *mongo.Database, config *Config.Config) *[]Models.Portfolio {
	collection := db.Collection("Portfolio")
	curr, err := collection.Find(context.TODO(), bson.M{})
	if (err != nil) {
		log.Println(err)
		return nil
	}

	var portfolios []Models.Portfolio
	for (curr.Next(context.TODO())) {
		elem := Models.Portfolio{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Println(errr)
			return nil
		}

		portfolios = append(portfolios, elem)
	}

	return &portfolios
}

func EditPortfolio(db *mongo.Database, config *Config.Config, updatedModel *Models.Portfolio) *Models.Portfolio {
	collection := db.Collection("Portfolio")
	exists := GetPortfolio(db, config, updatedModel.Id.Hex())
	if (exists == nil) {
		return nil
	}

	res := collection.FindOneAndReplace(context.TODO(), bson.M{"_id": updatedModel.Id}, updatedModel)
	if (res.Err() != nil) {
		log.Println(res.Err())
		return nil
	}

	result := GetPortfolio(db, config, updatedModel.Id.Hex())
	if (result == nil) {
		return nil
	}

	return result
}

func DeletePortfolio(db *mongo.Database, config *Config.Config, id string) bool {
	collection := db.Collection("Portfolio")
	_id, iderr := primitive.ObjectIDFromHex(id)
	if (iderr != nil) {
		log.Println(iderr)
		return false;
	}

	res := GetPortfolio(db, config, id)
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