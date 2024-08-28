package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Client
}

func NewDatabase(dsn string) (*Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{db: client}, nil
}

func (d *Database) Close() error {
	return d.db.Disconnect(context.TODO())
}

func (d *Database) GetDB() *mongo.Client {
	return d.db
}
