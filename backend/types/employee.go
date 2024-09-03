package types

type EmployeeStore interface {
	CreateEmployee(Employee) error
	GetEmployee(int) (*Employee, error)
	GetEmployees() ([]Employee, error)
	// UpdateEmployee(Employee) error
	DeleteEmployee(int) error
}

type Employee struct {
	ID int `bson:"_id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName string `bson:"last_name" json:"last_name"`
	BirthDate string `bson:"birth_date" json:"birth_date"`
	Telephone string `bson:"telephone" json:"telephone"`
	Email string `bson:"email" json:"email"`
	Address string `bson:"address" json:"address"`
}