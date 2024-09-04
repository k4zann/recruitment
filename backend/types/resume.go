package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResumeStore interface {
	CreateResume(Resume) error
	GetResume(primitive.ObjectID) (*Resume, error)
	GetResumes() ([]Resume, error)
	DeleteResume(primitive.ObjectID) error
}

type Resume struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EmployeeID     primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	Education      Education          `bson:"education" json:"education"`
	Experience     []Experience       `bson:"experience" json:"experience"`
	Skills         []Skill            `bson:"skills" json:"skills"`
	DisabilityInfo *DisabilityInfo    `bson:"disability_info" json:"disability_info"`
	AdditionalInfo *AdditionalInfo    `bson:"additional_info" json:"additional_info"`
	JobExpectation JobExpectation     `bson:"job_expectation" json:"job_expectation"`
}

type Education struct {
	InstitutionName string   `json:"institution_name"`
	Location        string   `json:"location"`
	Specialization  string   `json:"specialization"`
	YearsAttended   string   `json:"years_attended"`
	Certificates    []string `json:"certificates,omitempty"`
}

type Experience struct {
	CompanyName      string `json:"company_name"`
	Position         string `json:"position"`
	WorkPeriod       string `json:"work_period"`
	Responsibilities string `json:"responsibilities"`
}

type Skill struct {
	ProfessionalSkills []string              `json:"professional_skills"`
	PersonalQualities  []string              `json:"personal_qualities"`
	Languages          []LanguageProficiency `json:"languages"`
}

type LanguageProficiency struct {
	Language string `json:"language"`
	Level    string `json:"level"`
}

type DisabilityInfo struct {
	DisabilityGroup     string `json:"disability_group"`
	DisabilityDetails   string `json:"disability_details,omitempty"`
	WorkConditions      string `json:"work_conditions,omitempty"`
	MedicalRestrictions string `json:"medical_restrictions,omitempty"`
}

type AdditionalInfo struct {
	RelocationWillingness bool   `json:"relocation_willingness"`
	TravelWillingness     bool   `json:"travel_willingness"`
	Hobbies               string `json:"hobbies,omitempty"`
}

type JobExpectation struct {
	DesiredPosition string `json:"desired_position"`
	ExpectedSalary  string `json:"expected_salary"`
}

type CreateResumePayload struct {
	EmployeeID     primitive.ObjectID `json:"employee_id" validate:"required"`
	Education      Education          `json:"education" validate:"required"`
	Experience     []Experience       `json:"experience" validate:"required"`
	Skills         []Skill            `json:"skills" validate:"required"`
	DisabilityInfo *DisabilityInfo    `json:"disability_info"`
	AdditionalInfo *AdditionalInfo    `json:"additional_info"`
	JobExpectation JobExpectation     `json:"job_expectation" validate:"required"`
}
