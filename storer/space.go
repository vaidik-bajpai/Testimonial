package storer

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storer) CreateSpace(ctx context.Context, space *Space) (primitive.ObjectID, error) {
	var err error
	space.ID = primitive.NewObjectID()
	space.CreatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		return primitive.NilObjectID, err
	}

	_, err = s.collections.Space.InsertOne(ctx, space)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return space.ID, nil
}

func (s *Storer) UpdateSpace(ctx context.Context, filter bson.M, updates bson.D) error {
	_, err := s.collections.Space.UpdateByID(ctx, bson.M{"_id": filter}, updates)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storer) DeleteSpace(ctx context.Context, id primitive.ObjectID) error {
	_, err := s.collections.Space.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (s *Storer) GetSpace(ctx context.Context, space *GetSpaceRes) error {
	err := s.collections.Space.FindOne(ctx, bson.M{"_id": space.ID}).Decode(&space)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storer) ListSpace(ctx context.Context) (*mongo.Cursor, error) {
	findOpts := options.Find().SetProjection(bson.M{
		"_id":   1,
		"title": 1,
		"logo":  1,
	})

	cur, err := s.collections.Space.Find(ctx, bson.M{}, findOpts)
	if err != nil {
		return nil, err
	}

	return cur, err
}
