package database

import (
	"context"

	"github.com/janapc/grpc-movies-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieDatabase struct {
	DB *mongo.Collection
}

func NewMovieDatabase(db *mongo.Collection) *MovieDatabase {
	return &MovieDatabase{
		DB: db,
	}
}

func (m *MovieDatabase) Save(movie *entity.Movie) (string, error) {
	result, err := m.DB.InsertOne(context.TODO(), movie)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *MovieDatabase) Update(movie *entity.Movie) error {
	filter := bson.D{{Key: "_id", Value: movie.ID}}
	update := bson.D{primitive.E{Key: "$set", Value: movie}}
	_, err := m.DB.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MovieDatabase) FindById(id string) (*entity.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	var movie *entity.Movie
	err = m.DB.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *MovieDatabase) FindAll() ([]entity.Movie, error) {
	filter := bson.D{{}}
	cursor, err := m.DB.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var movies []entity.Movie
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}
	return movies, err
}

func (m *MovieDatabase) Remove(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = m.DB.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
