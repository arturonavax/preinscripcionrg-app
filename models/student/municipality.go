package student

import "time"

//Municipality Municipio de la aplicacion.
type Municipality struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID  int    `json:"userID" gorm:"type:integer"`
	StateID int    `json:"stateID" gorm:"type:integer"`
	Name    string `json:"name" gorm:"not null;unique;type:varchar(50)"`
}
