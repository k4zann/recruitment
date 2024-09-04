package vacancy

import (
	"context"
	"recruitment/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	MongoCollection *mongo.Collection
}

func NewStore(mongoCollection *mongo.Collection) *Store {
	return &Store{MongoCollection: mongoCollection}
}

func (s *Store) CreateVacancy(vacancy types.Vacancy) error {
	_, err := s.MongoCollection.InsertOne(
		context.Background(),
		vacancy,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetVacancy(id primitive.ObjectID) (*types.Vacancy, error) {
	var vacancy types.Vacancy

	err := s.MongoCollection.FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	).Decode(&vacancy)
	if err != nil {
		return nil, err
	}

	return &vacancy, nil
}

func (s *Store) GetVacancies() ([]types.Vacancy, error) {
	var vacancies []types.Vacancy

	cursor, err := s.MongoCollection.Find(
		context.Background(),
		bson.M{},
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var vacancy types.Vacancy

		err := cursor.Decode(&vacancy)
		if err != nil {
			return nil, err
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (s *Store) DeleteVacancy(id primitive.ObjectID) error {
	_, err := s.MongoCollection.DeleteOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
