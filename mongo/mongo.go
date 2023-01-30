package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mongo/collection"
)

type database struct {
	db     *mongo.Database
	client *mongo.Client
}

func ConnectDB(connect Connect, dbName string) (database, error) {
	db, client, err := connect.Connect(dbName)
	if err != nil {
		return database{}, err
	}

	return database{db: db, client: client}, nil
}

func (d database) Collection(name string) Collection {
	return collection.New(d.db, d.client, name)
}
