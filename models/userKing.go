package models

import "time"

//UserKind Tipo de Usuario de la Aplicacion
type UserKind struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID      int    `json:"userID" gorm:"type:integer"`
	Name        string `json:"name" gorm:"not null;unique;type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(200)"`

	// Permisos sobre Usuarios
	ReadUsers   bool `json:"readUsers" gorm:"not null;type:boolean;default:false"`
	UpdateUsers bool `json:"updateUsers" gorm:"not null;type:boolean;default:false"`
	DeleteUsers bool `json:"deleteUsers" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre tu usuario
	UpdateMe bool `json:"updateMe" gorm:"not null;type:boolean;default:true"`

	// Permisos sobre Estudiantes
	ReadStudents   bool `json:"readStudent" gorm:"not null;type:boolean;default:false"`
	CreateStudents bool `json:"createStudent" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Asignaturas
	ReadAsignatures bool `json:"readAsignatures" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Paises
	ReadCountries bool `json:"readCountries" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Menciones
	ReadMentions bool `json:"readMentions" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Municipios
	ReadMunicipalities bool `json:"readMunicipalities" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Parroquias
	ReadParishes bool `json:"readParishes" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Secciones
	ReadSections bool `json:"readSections" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Sectores
	ReadSectors bool `json:"readSectors" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Estados
	ReadStates bool `json:"readStates" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Profesores
	ReadTeachers bool `json:"readTeachers" gorm:"not null;type:boolean;default:false"`

	// Permisos sobre Tipos de vias
	ReadTypeOfRoads bool `json:"readTypeOfRoads" gorm:"not null;type:boolean;default:false"`
}
