package storer

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Storer struct {
	db          *mongo.Client
	collections *Collections
}

type Collections struct {
	Space       *mongo.Collection
	Testimonial *mongo.Collection
}

func NewStorer(db *mongo.Client, collections *Collections) *Storer {
	return &Storer{
		db:          db,
		collections: collections,
	}
}

func MakeCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("testimonials").Collection(collectionName)
	return collection
}
