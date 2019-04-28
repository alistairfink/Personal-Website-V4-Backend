package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Education struct
{
	Id primitive.ObjectID `json:"id" bson:"_id"`
	School string `json"school" bson:"school"`
	Start string `json"start" bson:"start"`
	End string `json"end" bson:"end"`
	Scholarship []string `json"scholarships" bson:"scholarships"`
	Award []string `json"awards" bson:"awards"`
	Location string `json"location" bson:"location"`
	Title string `json"title" bson:"title"`
	NotableProject []string `json"notableProj" bson:"notableProj"`
	ExtraCurcicular []string `json"extraCuric" bson:"extraCuric"`
	Image string `json"img" bson:"img`	
}