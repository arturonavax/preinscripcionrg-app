package student

import "time"

//Student Estudiante de la institucion
type Student struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	//IDs
	UserID int `json:"userId" gorm:"not null;type:integer"`

	CountryOfBirthID      int `json:"countryOfBirthID" gorm:"not null;type:integer"`
	StateOfBirthID        int `json:"stateOfBirthID" gorm:"not null;type:integer"`
	MunicOfBirthID        int `json:"municOfBirthID" gorm:"not null;type:integer"`
	MunicipalityID        int `json:"municipalityID" gorm:"not null;type:integer"`
	InstitutionOfOriginID int `json:"institutionOfOriginID" gorm:"not null;type:integer"`
	ParishID              int `json:"parishID" gorm:"not null;type:integer"`
	SectorID              int `json:"sectorID" gorm:"not null;type:integer"`
	TypeOfRoadID          int `json:"typeOfRoadID" gorm:"not null;type:integer"`
	FatherID              int `json:"fatherID" gorm:"not null;type:integer"`
	MotherID              int `json:"motherID" gorm:"not null;type:integer"`
	MentionID             int `json:"mentionID" gorm:"not null;type:integer"`
	SectionID             int `json:"sectionID" gorm:"not null;type:integer"`
	StudentConditionID    int `json:"studentConditionID" gorm:"not null; type:integer"`
	RepresentativeID      int `json:"representativeID" gorm:"not null; type:integer"`
	TeacherID             int `json:"teacherID" gorm:"not null; type:integer"`
	//-----

	// Datos del Estudiante.
	SIGECOD string `json:"SIGECOD" gorm:"unique;type:varchar(20)"`

	MotherFirstName   string `json:"motherFirstName" gorm:"-"`
	MotherLastName    string `json:"motherLastName" gorm:"-"`
	MotherEmail       string `json:"motherEmail" gorm:"-"`
	MotherPhoneNumber string `json:"motherPhoneNumber" gorm:"-"`
	MotherCIType      string `json:"motherCIType" gorm:"-"`
	MotherCI          int    `json:"motherCI" gorm:"-"`

	FatherFirstName   string `json:"fatherFirstName" gorm:"-"`
	FatherLastName    string `json:"fatherLastName" gorm:"-"`
	FatherEmail       string `json:"fatherEmail" gorm:"-"`
	FatherPhoneNumber string `json:"fatherPhoneNumber" gorm:"-"`
	FatherCIType      string `json:"fatherCIType" gorm:"-"`
	FatherCI          int    `json:"fatherCI" gorm:"-"`

	RepresentativeFirstName   string `json:"representativeFirstName" gorm:"-"`
	RepresentativeLastName    string `json:"representativeLastName" gorm:"-"`
	RepresentativeEmail       string `json:"representativeEmail" gorm:"-"`
	RepresentativePhoneNumber string `json:"representativePhoneNumber" gorm:"-"`
	RepresentativeCIType      string `json:"representativeCIType" gorm:"-"`
	RepresentativeCI          int    `json:"representativeCI" gorm:"-"`

	FirstName string `json:"firstName" gorm:"not null;type:varchar(50)"`
	LastName  string `json:"lastName" gorm:"not null;type:varchar(50)"`

	CIType string `json:"ciType" gorm:"not null;type:char(1)"`
	CI     int    `json:"ci" gorm:"not null; unique;integer"`

	DateOfBirth string `json:"dateOfBirth" gorm:"not null;type:varchar(20)"`

	Gender string `json:"gender" gorm:"not null;type:char(1)"`

	HealthProblem  bool   `json:"healthProblem" gorm:"not null;type:boolean"`
	HealthProblemE string `json:"healthProblemE" gorm:"type:varchar(300)"`

	Email       string `json:"email" gorm:"not null;type:varchar(30)"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(30)"`

	Address     string `json:"address" gorm:"not null;type:varchar(30)"`
	Scholarship bool   `json:"scholarship" gorm:"not null;type:boolean"`
	Canaima     bool   `json:"canaima" gorm:"not null;type:boolean"`

	ConditionOfHousing string `json:"conditionOfHousing" gorm:"not null;type:varchar(30)"`

	//Escolaridad-
	Year int `json:"year" gorm:"not null;type:integer"`

	Age    int    `json:"age" gorm:"not null;type:integer"`
	Size   string `json:"size" gorm:";type:varchar(10)"`
	Weight int    `json:"weight" gorm:";type:integer"`

	RepeatAsignature  string `json:"repeatAsignature" gorm:"type:varchar(200)"`
	PendingAsignature string `json:"pendingAsignature" gorm:"type:varchar(200)"`
	Regular           bool   `json:"regular" gorm:"not null;type:boolean"`
	InscriptionDate   string `json:"inscriptionDate" gorm:"not null;type:varchar(20)"`

	Message string `json:"message" gorm:"-"`
	Code    int    `json:"code" gorm:"-"`
}
