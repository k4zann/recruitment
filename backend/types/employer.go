package types

type EmployerStore interface {
	CreateEmployer(Employer) error
	GetEmployer(int) (*Employer, error)
	GetEmployers() ([]Employer, error)
	// UpdateEmployer(Employer) error
	DeleteEmployer(int) error
}

type Employer struct {
	ID int `bson:"_id" json:"id"`
	CompanyName string `bson:"company_name" json:"company_name"`
	Sphere string `bson:"sphere" json:"sphere"`
	Address string `bson:"address" json:"address"`
	Telephone string `bson:"telephone" json:"telephone"`
	Email string `bson:"email" json:"email"`
	NumberOfEmployees int `bson:"number_of_employees" json:"number_of_employees"`
}

