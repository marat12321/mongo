package filter

import "go.mongodb.org/mongo-driver/mongo"

type Lookup struct {
	From         string         `bson:"from,omitempty"`
	As           string         `bson:"as,omitempty"`
	DB           string         `bson:"DB,omitempty"`
	LocalField   string         `bson:"localField,omitempty"`
	ForeignField string         `bson:"foreignField,omitempty"`
	Pipeline     mongo.Pipeline `bson:"pipeline,omitempty"`
}
