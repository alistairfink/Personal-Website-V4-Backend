package Models

type CreateExperience struct
{
	Position string `json"position" bson:"position"`
	Start string `json"start" bson:"start"`
	End string `json"end" bson:"end"`
	Company string `json"comp" bson:"comp"`
	Location string `json"location" bson:"location"`
	Data []string `json"data" bson:"data"`
	Image string `json"img" bson:"img"`
}