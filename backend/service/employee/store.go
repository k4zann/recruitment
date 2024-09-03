package employee

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

func (s *Store) CreateEmployee(employee types.Employee) error {
	_, err := s.MongoCollection.InsertOne(context.Background(), employee)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetEmployee(id int) (*types.Employee, error) {
	var employee types.Employee

	err := s.MongoCollection.FindOne(
		context.Background(), 
		map[string]int{"_id": id}).Decode(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (s *Store) GetEmployees() ([]types.Employee, error) {
	var employees []types.Employee

	cursor, err := s.MongoCollection.Find(
		context.Background(), 
		map[string]int{},
	)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var employee types.Employee
	
		err := cursor.Decode(&employee)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (s *Store) DeleteEmployee(id int) error {
	_, err := s.MongoCollection.DeleteOne(
		context.Background(), 
		map[string]int{"_id": id},
	)
	if err != nil {
		return err
	}

	return nil
}