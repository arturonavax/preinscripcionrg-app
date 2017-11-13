package models

import "time"

//People Persona de la aplicacion
type People struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	FirstName   string `json:"firstName" gorm:"not null;type:varchar(30)"`
	LastName    string `json:"lastName" gorm:"not null;type:varchar(100)"`
	Email       string `json:"email" gorm:"not null; unique;type:varchar(100)"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(20)"`
	Address     string `json:"address" gorm:"type:varchar(200)"`
}
