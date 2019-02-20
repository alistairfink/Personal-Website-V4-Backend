package Controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/alistairfink/Personal-Website-V4-Backend/src/config"
)

type Controller struct {
	DB *mongo.Client
	Config *Config.Config
}