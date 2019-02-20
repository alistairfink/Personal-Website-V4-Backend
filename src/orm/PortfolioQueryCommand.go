package Orm

import (
	"context"
	"log"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/alistairfink/Personal-Website-V4-Backend/src/models"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

func AddPortfolio(db *mongo.Database, config *Config.Config) *Models.Portfolio {
	return nil
}

func GetPortfolio(db *mongo.Database, config *Config.Config) *Models.Portfolio {
	return nil
}

func GetAllPortfolio(db *mongo.Database, config *Config.Config) *[]Models.Portfolio {
	collection := db.Collection("Portfolio")
	curr, err := collection.Find(context.TODO(), bson.D{})
	if (err != nil) {
		log.Fatal(err)
	}

	var portfolios []Models.Portfolio
	for (curr.Next(context.TODO())) {
		elem := Models.Portfolio{}
		errr := curr.Decode(&elem)
		if (errr != nil) {
			log.Fatal(errr)
		}

		portfolios = append(portfolios, elem)
	}

	return &portfolios
}

func EditPortfolio(db *mongo.Database, config *Config.Config) *Models.Portfolio {
	return nil
}

func DeletePortfolio(db *mongo.Database, config *Config.Config) bool {
	return true
}