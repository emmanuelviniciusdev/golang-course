package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieNetflix struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	MovieName string `json:"movieName"`
	Watched bool `json:"watched"`
}
