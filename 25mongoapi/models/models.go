package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieNetflixInput struct {
	MovieName string `json:"movieName" bson:"movieName"`
	Watched bool `json:"watched" bson:"watched"`
}

type MovieNetflix struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	MovieName string `json:"movieName" bson:"movieName"`
	Watched bool `json:"watched" bson:"watched"`
}
