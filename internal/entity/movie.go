package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	ImageUrl    string             `bson:"image_url"`
}

func NewMovie(title, description, imageUrl string) *Movie {
	return &Movie{
		Title:       title,
		Description: description,
		ImageUrl:    imageUrl,
	}
}
