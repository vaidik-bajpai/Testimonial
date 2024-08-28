package storer

import "go.mongodb.org/mongo-driver/mongo"

type Storer struct {
	db *mongo.Client
}

func NewStorer(db *mongo.Client) *Storer {
	return &Storer{
		db: db,
	}
}
