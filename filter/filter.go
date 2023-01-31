package filter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type filter struct {
	Filter []bson.D
}

// Set изменяет значение поля/полей при их получении на установленное в фильтре
func (f *filter) Set(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$set", val}})
	return f
}

// Project спецификация полей, которые будут отображены
func (f *filter) Project(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$project", val}})
	return f
}

// AddFields спецификация дополнительных полей, которые должны быть получены в результате запроса
func (f *filter) AddFields(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$addFields", val}})
	return f
}

// Match спецификация для поиска записей согласно заданному фильтру
func (f *filter) Match(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$match", val}})
	return f
}

// Lookup присоединение таблиц к результату запроса
func (f *filter) Lookup(val Lookup) *filter {
	f.Filter = append(f.Filter, bson.D{{"$lookup", val}})
	return f
}

// Skip skip
func (f *filter) Skip(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$skip", val}})
	return f
}

// Sort sorting
func (f *filter) Sort(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$sort", val}})
	return f
}

// Unwind разделение поля выводящего массив на несколько результатов
func (f *filter) Unwind(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$unwind", val}})
	return f
}

func (f *filter) Unset(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$unset", val}})
	return f
}

// Limit limit
func (f *filter) Limit(val int64) *filter {
	f.Filter = append(f.Filter, bson.D{{"$limit", val}})
	return f
}

// Group grouping
func (f *filter) Group(val interface{}) *filter {
	f.Filter = append(f.Filter, bson.D{{"$group", val}})
	return f
}

// Use convert filter to mongo.Pipeline
func (f *filter) Use() mongo.Pipeline {
	return f.Filter
}
