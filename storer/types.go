package storer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Space struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" form:"_"`
	Name              string             `bson:"name" form:"name" binding:"required,min=3,max=30"`
	Logo              string             `bson:"logo" form:"_"`
	Title             string             `bson:"title" form:"title" binding:"required,min=3,max=30"`
	Message           string             `bson:"message" form:"message" binding:"required,min=10,max=200"`
	Questions         []string           `bson:"questions" form:"questions" binding:"required,atLeastOne"`
	CollectInfo       []string           `bson:"collect_info" form:"collect_info" binding:"required,atLeastName"`
	CollectionType    string             `bson:"collection_type" form:"collection_type" binding:"required,typesAllowed"`
	CollectStarRating bool               `bson:"collect_star_rating" form:"collect_star_rating" binding:"required"`
	CustomButtonColor string             `bson:"custom_button_color" form:"custom_button_color" binding:"required" validate:"hexcolor"`
	Language          string             `bson:"language" form:"language" binding:"required"`
	CreatedAt         time.Time          `bson:"created_at" form:"_"`
	UpdatedAt         time.Time          `bson:"updated_at" form:"_"`
}

type UpdateSpaceRequest struct {
	ID                *primitive.ObjectID `bson:"_id" form:"-"`
	Name              *string             `bson:"name" form:"name" binding:"omitempty,min=3,max=30"`
	Logo              *string             `bson:"logo" form:"_"`
	Title             *string             `bson:"title" form:"title" binding:"omitempty,min=3,max=30"`
	Message           *string             `bson:"message" form:"message" binding:"omitempty,min=10,max=200"`
	Questions         []string            `bson:"questions" form:"questions" binding:"omitempty,atLeastOne"`
	CollectInfo       []string            `bson:"collect_info" form:"collect_info" binding:"omitempty,atLeastName"`
	CollectionType    *string             `bson:"collection_type" form:"collection_type" binding:"omitempty,typesAllowed"`
	CollectStarRating *bool               `bson:"collect_star_rating" form:"collect_star_rating" binding:"omitempty"`
	CustomButtonColor *string             `bson:"custom_button_color" form:"custom_button_color" binding:"omitempty" validate:"omitempty,hexcolor"`
	Language          *string             `bson:"language" form:"language" binding:"omitempty"`
	UpdatedAt         *time.Time          `bson:"updated_at" form:"-"`
}

type ListSpaceRes struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
	Logo  string             `bson:"logo" json:"logo"`
}

type GetSpaceRes struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	Name              string             `bson:"name" json:"name"`
	Logo              string             `bson:"logo" json:"logo"`
	Title             string             `bson:"title" json:"title"`
	Message           string             `bson:"message" json:"message"`
	Questions         []string           `bson:"questions" json:"questions"`
	CollectInfo       []string           `bson:"collect_info" json:"collect_info"`
	CollectionType    string             `bson:"collection_type" json:"collection_type"`
	CollectStarRating bool               `bson:"collect_star_rating" json:"collect_star_rating"`
	CustomButtonColor string             `bson:"custom_button_color" json:"custom_button_color"`
	Language          string             `bson:"language" json:"language"`
	CreatedAt         time.Time          `bson:"_" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}
