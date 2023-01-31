package mongo

import (
	. "github.com/marat12321/mongo/collection"
	. "github.com/marat12321/mongo/connectOptions"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	db     *mongo.Database
	client *mongo.Client
}

func ConnectDB(connect Opt, dbName string) (database, error) {
	db, client, err := connect.Connect(dbName)
	if err != nil {
		return database{}, err
	}

	return database{db: db, client: client}, nil
}

func (d database) Collection(name string) Collection {
	return New(d.db, d.client, name)
}
