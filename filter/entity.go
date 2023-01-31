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

type QueryFilter interface {
	Set(val interface{}) *filter
	Project(val interface{}) *filter
	AddFields(val interface{}) *filter
	Match(val interface{}) *filter
	Lookup(val Lookup) *filter
	Skip(val interface{}) *filter
	Sort(val interface{}) *filter
	Unwind(val interface{}) *filter
	Unset(val interface{}) *filter
	Limit(val int64) *filter
	Group(val interface{}) *filter
	Use() mongo.Pipeline
}
