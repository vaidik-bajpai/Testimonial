package storer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Space struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name              string             `bson:"name" json:"name" binding:"required,min=3,max=30"`
	Logo              string             `bson:"logo" json:"logo" binding:"required"`
	Title             string             `bson:"title" json:"title" binding:"required,min=3,max=30"`
	Message           string             `bson:"message" json:"message" binding:"required,min=10,max=200"`
	Questions         []string           `bson:"questions" json:"questions" binding:"min=1,max=8,dive,required,min=10,max=30"`
	CollectInfo       []Field            `bson:"collect_info" json:"collect_info" binding:"required,atLeastName,dive"`
	CollectionType    string             `bson:"collection_type" json:"collection_type" binding:"required,typesAllowed"`
	CollectStarRating bool               `bson:"collect_star_rating" json:"collect_star_rating" binding:"required"`
	CustomButtonColor string             `bson:"custom_button_color" json:"custom_button_color" binding:"required" validate:"hexcolor"`
	Language          string             `bson:"language" json:"language" binding:"required"`
	CreatedAt         time.Time          `bson:"created_at" json:"-"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"-"`
}

type UpdateSpaceReq struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Name              string             `bson:"name" json:"name" binding:"min=3,max=30"`
	Logo              string             `bson:"logo" json:"logo"`
	Title             string             `bson:"title" json:"title" binding:"min=3,max=30"`
	Message           string             `bson:"message" json:"message" binding:"min=10,max=200"`
	Questions         []string           `bson:"questions" json:"questions" binding:"min=1,max=8,dive,required,min=10,max=30"`
	CollectInfo       []Field            `bson:"collect_info" json:"collect_info" binding:"required,atLeastName,dive"`
	CollectionType    string             `bson:"collection_type" json:"collection_type" binding:"typesAllowed"`
	CollectStarRating bool               `bson:"collect_star_rating" json:"collect_star_rating"`
	CustomButtonColor string             `bson:"custom_button_color" json:"custom_button_color"`
	Language          string             `bson:"language" json:"language"`
	CreatedAt         time.Time          `bson:"created_at" json:"-"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"-"`
}

type ListSpaceRes struct {
	ID    primitive.ObjectID `bson:"_id" json:"id" binding:"required"`
	Title string             `bson:"title" json:"title" binding:"required"`
	Logo  string             `bson:"logo" json:"logo" binding:"required"`
}

type GetSpaceRes struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	Name              string             `bson:"name" json:"name"`
	Logo              string             `bson:"logo" json:"logo"`
	Title             string             `bson:"title" json:"title"`
	Message           string             `bson:"message" json:"message"`
	Questions         []string           `bson:"questions" json:"questions"`
	CollectInfo       []Field            `bson:"collect_info" json:"collect_info" binding:"required,atLeastName,dive"`
	CollectionType    string             `bson:"collection_type" json:"collection_type"`
	CollectStarRating bool               `bson:"collect_star_rating" json:"collect_star_rating"`
	CustomButtonColor string             `bson:"custom_button_color" json:"custom_button_color"`
	Language          string             `bson:"language" json:"language"`
	CreatedAt         time.Time          `bson:"_" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}

type TextTestimonial struct {
	ID        primitive.ObjectID     `bson:"_id" json:"-"`
	Type      string                 `bson:"type" json:"type" binding:"required"`
	Message   string                 `bson:"message" json:"message" binding:"required"`
	UserInfo  map[string]interface{} `bson:"user_info" json:"user_info" binding:"required"`
	Image     string                 `bson:"image" json:"image"`
	Avatar    string                 `bson:"avatar" json:"avatar"`
	Rating    int64                  `bson:"rating" json:"rating" binding:"gte=0,lte=5"`
	SpaceID   primitive.ObjectID     `bson:""`
	CreatedAt time.Time              `bson:"created_at" json:"-"`
	UpdatedAt time.Time              `bson:"updated_at" json:"-"`
}

type VideoTestimonial struct {
	ID        primitive.ObjectID     `bson:"_id" json:"-"`
	Type      string                 `bson:"type" json:"type"`
	UserInfo  map[string]interface{} `bson:"user_info" json:"user_info" binding:"required"`
	Video     string                 `bson:"video" json:"video"`
	Rating    int64                  `bson:"rating" json:"rating" binding:"gte=0,lte=5"`
	SpaceID   primitive.ObjectID     `bson:"space_id" uri:""`
	CreatedAt time.Time              `bson:"created_at" json:"-"`
	UpdatedAt time.Time              `bson:"updated_at" json:"-"`
}

type GetTestimonial struct {
	ID          primitive.ObjectID `bson:"_id" json:"-"`
	Type        string             `bson:"type" json:"type" binding:"required"`
	Message     string             `bson:"message" json:"message" binding:"required"`
	UserDetails interface{}        `bson:"user_details" json:"user_details" binding:"required,dive"`
	Image       string             `bson:"image" json:"image,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Video       string             `bson:"video" json:"video,omitempty"`
	Rating      int64              `bson:"rating" json:"rating" binding:"gte=0,lte=5"`
	SpaceID     primitive.ObjectID `bson:""`
	CreatedAt   time.Time          `bson:"created_at" json:"-"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"-"`
}

type Field struct {
	FieldName  string `bson:"field_name" json:"field_name" binding:"required"`
	IsRequired bool   `bson:"is_required" json:"is_required" binding:"required"`
	IsVisible  bool   `bson:"is_visible" json:"is_visible" binding:"required"`
}
