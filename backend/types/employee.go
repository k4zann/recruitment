package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeStore interface {
	CreateEmployee(Employee) error
	GetEmployee(primitive.ObjectID) (*Employee, error)
	GetEmployees() ([]Employee, error)
	// UpdateEmployee(Employee) error
	DeleteEmployee(primitive.ObjectID) error
}

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	BirthDate string             `bson:"birth_date" json:"birth_date"`
	Telephone string             `bson:"telephone" json:"telephone"`
	Email     string             `bson:"email" json:"email"`
	Address   string             `bson:"address" json:"address"`
}

type CreateEmployeePayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}
