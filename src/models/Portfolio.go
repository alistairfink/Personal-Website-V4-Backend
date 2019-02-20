package Models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type Portfolio struct
{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	DisplayImage string `json:"dspImg" bson:"dspImg"`
	Description []string `json:"desc" bson:"desc"`
	Image []string `json:"images" bson:"images"`
	Video []string `json:"videos" bson:"videos"`
	Year string `json:"year" bson:"year"`
}