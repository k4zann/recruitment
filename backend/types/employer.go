package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployerStore interface {
	CreateEmployer(Employer) error
	GetEmployer(primitive.ObjectID) (*Employer, error)
	GetEmployers() ([]Employer, error)
	// UpdateEmployer(Employer) error
	DeleteEmployer(primitive.ObjectID) error
}

type Employer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CompanyName       string             `bson:"company_name" json:"company_name"`
	Sphere            string             `bson:"sphere" json:"sphere"`
	Address           string             `bson:"address" json:"address"`
	Telephone         string             `bson:"telephone" json:"telephone"`
	Email             string             `bson:"email" json:"email"`
	NumberOfEmployees int                `bson:"number_of_employees" json:"number_of_employees"`
}

type CreateEmployerPayload struct {
	CompanyName       string `json:"company_name"`
	Sphere            string `json:"sphere"`
	Address           string `json:"address"`
	Telephone         string `json:"telephone"`
	Email             string `json:"email"`
	NumberOfEmployees int    `json:"number_of_employees"`
}
