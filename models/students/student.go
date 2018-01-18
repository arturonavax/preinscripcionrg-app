package student

import "time"

//Student Estudiante de la institucion
type Student struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	//IDs
	UserID uint `json:"userId" gorm:"not null;type:integer"`

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
	SectionsID            int `json:"sectionID" gorm:"not null;type:integer"`
	StudentConditionID    int `json:"studentConditionID" gorm:"not null; type:integer"`
	RepresentativeID      int `json:"representativeID" gorm:"not null; type:integer"`
	TeacherID             int `json:"teacherID" gorm:"not null; type:integer"`
	//-----

	// Datos del Estudiante.
	SIGECOD string `json:"SIGECOD" gorm:"type:varchar(20)"`

	FirstName string `json:"firstName" gorm:"not null;type:varchar(50)"`
	LastName  string `json:"lastName" gorm:"not null;type:varchar(50)"`

	CIType string `json:"ciType" gorm:"not null;type:char(1)"`
	CI     int    `json:"ci" gorm:"not null; unique;integer"`

	DateOfBirth string `json:"dateOfBirth" gorm:"not null;type:date"`

	Gender string `json:"gender" gorm:"not null;type:char(1)"`

	HealthProblem  string `json:"healthProblem" gorm:"not null;type:boolean"`
	HealthProblemE string `json:"healthProblemE" gorm:"not null;type:varchar(300)"`

	Email       string `json:"email" gorm:"not null;type:varchar(30)"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null;type:varchar(12)"`

	Address     string `json:"address" gorm:"not null;type:varchar(30)"`
	Scholarship string `json:"scholarship" gorm:"not null;type:boolean"`
	Canaima     string `json:"canaima" gorm:"not null;type:boolean"`

	ConditionOfHousing string `json:"conditionOfHousing" gorm:"not null;type:varchar(30)"`

	//Escolaridad-
	Year int `json:"year" gorm:"not null;type:integer"`

	Age    int    `json:"age" gorm:"not null;type:integer"`
	Size   string `json:"size" gorm:"not null;type:varchar(10)"`
	Weight int    `json:"weight" gorm:"not null;type:integer"`

	RepeatAsignature  []Asignature `json:"repeatAsignature"`
	PendingAsignature []Asignature `json:"pendingAsignature"`
	Regular           bool         `json:"regular" gorm:"not null;type:boolean"`
	InscriptionDate   time.Time    `json:"inscriptionDate" gorm:"not null;type:date"`
}
