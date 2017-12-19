package student

import "time"

//Municipality Municipio de la aplicacion.
type Municipality struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID  uint   `json:"userID" gorm:"type:integer"`
	StateID uint   `json:"stateID" gorm:"type:integer"`
	Name    string `json:"name" gorm:"not null;unique;type:varchar(50)"`
}
