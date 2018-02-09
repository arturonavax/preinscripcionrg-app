package models

import (
	"time"
)

//User Usuario de la aplicacion
type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	PeopleID uint `json:"peopleID" gorm:"not null;unique;type:integer;"`
	StatusID uint `json:"statusID" gorm:"type:integer;default:1"`
	KindID   uint `json:"kindID" gorm:"type:integer;default:1"`

	Username string `json:"username" gorm:"not null;unique;type:varchar(30)"`
	Password string `json:"password" gorm:"not null;type:varchar(256)"`

	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"`

	FirstName   string `json:"firstName" gorm:"-"`
	LastName    string `json:"lastName" gorm:"-"`
	Email       string `json:"email" gorm:"-"`
	PhoneNumber string `json:"phoneNumber" gorm:"-"`
	Address     string `json:"address" gorm:"-"`
}
