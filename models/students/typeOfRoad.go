package student

import "time"

//TypeOfRoad Tipo de via de la aplicacion.
type TypeOfRoad struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID uint   `json:"userID" gorm:"type:integer"`
	Name   string `json:"name" gorm:"not null;unique;type:varchar(50)"`
}
