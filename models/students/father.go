package student

import "time"

//Father Padre de la aplicacion.
type Father struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID      uint   `json:"userID" gorm:"type:integer"`
	FirstName   string `json:"firstName" gorm:"not null;type:varchar(30)"`
	LastName    string `json:"lastName" gorm:"not null;type:varchar(30)"`
	Email       string `json:"email" gorm:"not null;unique;type:varchar(40)"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null;type:varchar(15)"`
	CI          string `json:"ci" gorm:"not null;unique;type:varchar(15)"`
}
