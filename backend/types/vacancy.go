package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type VacancyStore interface {
	CreateVacancy(Vacancy) error
	GetVacancy(primitive.ObjectID) (*Vacancy, error)
	GetVacancies() ([]Vacancy, error)
	DeleteVacancy(primitive.ObjectID) error
}

type Vacancy struct {
	ID                   primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	EmployerID           primitive.ObjectID    `bson:"employer_id" json:"employer_id"`
	Position             string                `json:"position"`
	Responsibilities     string              `json:"responsibilities"`
	QualificationReqs    []string              `json:"qualification_reqs"`
	WorkSchedule         string                `json:"work_schedule"`
	WorkingConditions    *WorkingConditions    `json:"working_conditions,omitempty"`
	SocialBenefits       *SocialBenefits       `json:"social_benefits,omitempty"`
	EmployeeExpectations *EmployeeExpectations `json:"employee_expectations,omitempty"`
}

type WorkingConditions struct {
	Accessibility     string `json:"accessibility"`
	Equipment         string `json:"equipment"`
	SpecialConditions string `json:"special_conditions"`
	Safety            string `json:"safety"`
}

type SocialBenefits struct {
	MedicalCare        string `json:"medical_care"`
	SocialCompensation string `json:"social_compensation"`
	ProfessionalGrowth string `json:"professional_growth"`
}

type EmployeeExpectations struct {
	TeamInteraction        string `json:"team_interaction"`
	QualityAndProductivity string `json:"quality_and_productivity"`
}

type CreateVacancyPayload struct {
	EmployerID           primitive.ObjectID    `json:"employer_id" validate:"required"`
	Position             string                `json:"position" validate:"required"`
	Responsibilities     string              `json:"responsibilities" validate:"required"`
	QualificationReqs    []string              `json:"qualification_reqs" validate:"required"`
	WorkSchedule         string                `json:"work_schedule" validate:"required"`
	WorkingConditions    *WorkingConditions    `json:"working_conditions,omitempty"`
	SocialBenefits       *SocialBenefits       `json:"social_benefits,omitempty"`
	EmployeeExpectations *EmployeeExpectations `json:"employee_expectations,omitempty"`
}
