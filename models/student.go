package models

import "time"

//Student Estudiante de la institucion
type Student struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID uint `json:"userId" gorm:"not null;type:integer"`

	// Datos del Estudiante.
	SIGECOD string `json:"SIGECOD" gorm:"type:varchar(20)"`

	FirstName string `json:"firstName" gorm:"not null;type:varchar(50)"`
	LastName  string `json:"lastName" gorm:"not null;type:varchar(50)"`

	CIType string `json:"ciType" gorm:"not null;type:char(1)"`
	CI     int    `json:"ci" gorm:"not null; unique;integer"`

	DateOfBirth    string `json:"dateOfBirth" gorm:"not null;type:date"`
	CountryOfBirth string `json:"countryOfBirth" gorm:"not null;type:varchar(30)"`
	StateOfBirth   string `json:"stateOfBirth" gorm:"not null;type:varchar(30)"`
	MunicOfBirth   string `json:"municOfBirth" gorm:"not null;type:varchar(30)"`

	Gender         string `json:"gender" gorm:"not null;type:char(1)"`
	CampusOfOrigin string `json:"campusOfOrigin" gorm:"not null;type:varchar(50)"`
	HealthProblem  string `json:"healthProblem" gorm:"not null;type:boolean"`
	HealthProblemE string `json:"healthProblemE" gorm:"not null;type:varchar(300)"`

	Email       string `json:"email" gorm:"not null;type:varchar(30)"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null;type:varchar(12)"`

	Address      string `json:"address" gorm:"not null;type:varchar(30)"`
	Municipality string `json:"municipality" gorm:"not null;type:varchar(30)"`
	Parish       string `json:"parish" gorm:"not null;type:varchar(30)"`
	Sector       string `json:"sector" gorm:"not null;type:varchar(30)"`
	TypeOfRoad   string `json:"typeOfRoad" gorm:"not null;type:varchar(30)"`
	Scholarship  string `json:"scholarship" gorm:"not null;type:boolean"`
	Canaima      string `json:"canaima" gorm:"not null;type:boolean"`

	ConditionOfHousing string `json:"conditionOfHousing" gorm:"not null;type:varchar(30)"`

	FirstNameM   string `json:"firstNameM" gorm:"not null;type:varchar(50)"`
	LastNameM    string `json:"lastNameM" gorm:"not null;type:varchar(50)"`
	PhoneNumberM string `json:"phoneNumberM" gorm:"not null;type:varchar(12)"`
	CIM          int    `json:"cim" gorm:"not null; unique;integer"`

	FirstNameP   string `json:"firstNameP" gorm:"not null;type:varchar(50)"`
	LastNameP    string `json:"lastNameP" gorm:"not null;type:varchar(50)"`
	PhoneNumberP string `json:"phoneNumberP" gorm:"not null;type:varchar(12)"`
	CIP          int    `json:"cip" gorm:"not null; unique;integer"`

	//Escolaridad
	Year    int    `json:"year" gorm:"not null;type:integer`
	Mention string `json:"mention" gorm:"not null;type:varchar(50)`
	Section string `json:"section" gorm:"not null; type:varchar(10)`

	Age    int    `json:"age" gorm:"not null;type:integer`
	Size   string `json:"size" gorm:"not null;type:varchar(10)`
	Weight int    `json:"weight" gorm:"not null;type:integer"`
}
