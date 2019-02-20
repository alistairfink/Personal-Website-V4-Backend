package Models

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type About struct
{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Description []string `json:"desc" bson:"desc"` 
	Image string `json:"img" bson:"img"`	
}