package employer

import (
	"context"
	"recruitment/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	MongoCollection *mongo.Collection
}

func NewStore(mongoCollection *mongo.Collection) *Store {
	return &Store{MongoCollection: mongoCollection}
}

func (s *Store) CreateEmployer(employer types.Employer) error {
	_, err := s.MongoCollection.InsertOne(context.Background(), employer)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetEmployer(id int) (*types.Employer, error) {
	var employer types.Employer

	err := s.MongoCollection.FindOne(
		context.Background(),
		map[string]int{"_id": id}).Decode(&employer)
	if err != nil {
		return nil, err
	}

	return &employer, nil
}

func (s *Store) GetEmployers() ([]types.Employer, error) {
	var employers []types.Employer

	cursor, err := s.MongoCollection.Find(
		context.Background(),
		map[string]int{},
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var employer types.Employer

		err := cursor.Decode(&employer)
		if err != nil {
			return nil, err
		}

		employers = append(employers, employer)
	}

	return employers, nil
}

func (s *Store) DeleteEmployer(id int) error {
	_, err := s.MongoCollection.DeleteOne(
		context.Background(),
		map[string]int{"_id": id},
	)
	if err != nil {
		return err
	}

	return nil
}