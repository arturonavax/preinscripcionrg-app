package student

import "time"

//State Estado de la aplicacion.
type State struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID    uint   `json:"userID" gorm:"type:integer"`
	CountryID uint   `json:"countryID" gorm:"type:integer"`
	Name      string `json:"name" gorm:"not null;unique;type:varchar(50)"`
}
