package connect

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type option struct {
	appName                *string
	auth                   *options.Credential
	hosts                  []string
	maxPoolSize            uint64
	minPoolSize            uint64
	poolMonitor            *event.PoolMonitor
	monitor                *event.CommandMonitor
	readConcern            *readconcern.ReadConcern
	replicaSet             *string
	retryReads             bool
	retryWrites            bool
	serverSelectionTimeout *time.Duration
	timeout                *time.Duration
	writeConcern           *writeconcern.WriteConcern
	zlibLevel              *int
	zstdLevel              *int
}

type Auth struct {
	DB       string
	UserName string
	Password string
}
