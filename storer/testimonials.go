package storer

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Storer) CreateTextTestimonial(ctx context.Context, tt TextTestimonial) error {
	_, err := s.collections.Testimonial.InsertOne(ctx, tt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storer) CreateVideoTestimonial(ctx context.Context, vt VideoTestimonial) error {
	_, err := s.collections.Testimonial.InsertOne(ctx, vt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storer) ListTestimonials(ctx context.Context, sID primitive.ObjectID) (interface{}, error) {
	cur, err := s.collections.Testimonial.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var testimonials []map[string]interface{}

	for cur.Next(context.Background()) {
		t := make(map[string]interface{})
		if err := cur.Decode(&t); err != nil {
			return nil, err
		}
		testimonials = append(testimonials, t)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return testimonials, err
}

func (s *Storer) GetTestimonial(ctx context.Context, tID, sID primitive.ObjectID) (GetTestimonial, error) {
	var testimonial GetTestimonial

	err := s.collections.Testimonial.FindOne(ctx, bson.M{"_id": tID, "spaceid": sID}).Decode(&testimonial)
	if err != nil {
		return GetTestimonial{}, err
	}

	return testimonial, nil
}

func (s *Storer) UpdateTestimonial(ctx context.Context) error {
	return nil
}

func (s *Storer) DeleteTestimonial(ctx context.Context, tID, sID primitive.ObjectID) error {
	_, err := s.collections.Testimonial.DeleteOne(ctx, bson.M{"_id": tID, "spaceid": sID})
	if err != nil {
		return err
	}

	return nil
}
