package Controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

type Controller struct {
	DB *mongo.Database
	Config *Config.Config
}