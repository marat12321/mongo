package mongo

import (
	"github.com/marat12321/mongo/collection"
	connect "github.com/marat12321/mongo/connectOptions"
	"github.com/marat12321/mongo/filter"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	db     *mongo.Database
	client *mongo.Client
}

func Transaction() chan collection.Transaction {
	return make(chan collection.Transaction, 1)
}

func ConnectDB(connect connect.Opt, dbName string) (database, error) {
	db, client, err := connect.Connect(dbName)
	if err != nil {
		return database{}, err
	}

	return database{db: db, client: client}, nil
}

func Connection() connect.Opt {
	return connect.Options()
}

func Filter() filter.QueryFilter {
	return filter.New()
}

func (d database) Collection(name string) Collection {
	return collection.New(d.db, d.client, name)
}
