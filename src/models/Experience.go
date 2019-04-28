package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct
{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Position string `json"position" bson:"position"`
	Start string `json"start" bson:"start"`
	End string `json"end" bson:"end"`
	Company string `json"comp" bson:"comp"`
	Location string `json"location" bson:"location"`
	Data []string `json"data" bson:"data"`
	Image string `json"img" bson:"img"`
}