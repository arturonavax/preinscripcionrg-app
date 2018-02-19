package student

import "time"

//Parish Parroquia de la aplicacion.
type Parish struct {
	ID        int `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserID         int    `json:"userID" gorm:"type:integer"`
	MunicipalityID int    `json:"municipalityID" gorm:"type:integer"`
	Name           string `json:"name" gorm:"not null;unique;type:varchar(50)"`

	Message string `json:"message" gorm:"-"`
	Code    int    `json:"code" gorm:"-"`
}
