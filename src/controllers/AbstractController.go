package Controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Controller struct {
	DB *mongo.Client
}