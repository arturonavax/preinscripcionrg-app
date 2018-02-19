package models

import "time"

//UserStatus Estatus del Usuario en la aplicacion
type UserStatus struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID      int    `json:"userID" gorm:"type:integer"`
	Name        string `json:"name" gorm:"not null;unique;varchar(50)"`
	Description string `json:"description" grom:"varchar(200)"`

	Message string `json:"message" gorm:"-"`
	Code    int    `json:"code" gorm:"-"`
}
