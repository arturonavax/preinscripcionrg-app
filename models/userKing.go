package models

import "time"

//UserKind Tipo de Usuario de la Aplicacion
type UserKind struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID      uint   `json:"userID" gorm:"type:integer"`
	Name        string `json:"name" gorm:"not null;unique;type:varchar(50)"`
	Description string `json:"description" gorm:"type:varchar(200)"`

	// Permisos sobre Usuarios
	ReadUsers   bool `json:"readUsers" gorm:"not null;type:boolean;default:false"`
	UpdateUsers bool `json:"updateUsers" gorm:"not null;type:boolean;default:false"`
	DeleteUsers bool `json:"deleteUsers" gorm:"not null;type:boolean;default:false"`
}
