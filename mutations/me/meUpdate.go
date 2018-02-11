package me

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

var MeUpdate = &graphql.Field{
	Type: graphqlTypes.UserType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"address": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		m := &models.Message{}
		peopleMe := &models.People{}

		peopleID := p.Context.Value("user").(models.User).PeopleID
		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.UpdateMe {
			db.Where("id = ?", peopleID).First(&peopleMe)

			_, isOK := p.Args["email"].(string)
			if isOK {
				peopleMe.Email, _ = p.Args["email"].(string)
			}

			_, isOK = p.Args["firstName"].(string)
			if isOK {
				peopleMe.FirstName, _ = p.Args["firstName"].(string)
			}

			_, isOK = p.Args["lastName"].(string)
			if isOK {
				peopleMe.LastName, _ = p.Args["lastName"].(string)
			}

			_, isOK = p.Args["phoneNumber"].(string)
			if isOK {
				peopleMe.PhoneNumber, _ = p.Args["phoneNumber"].(string)
			}

			_, isOK = p.Args["address"].(string)
			if isOK {
				peopleMe.Address, _ = p.Args["address"].(string)
			}

			err := db.Save(&peopleMe).Error
			if err != nil {
				m.Message = "#UserMe# : Ocurrio un Error."
				m.Code = http.StatusInternalServerError
			} else {
				m.Message = "#UserMe# : Usuario actualizado con exito"
				m.Code = http.StatusContinue
			}
		} else {
			m.Code = http.StatusNonAuthoritativeInfo
			m.Message = "#UserMe# : No tienes permisos para actualizar tu usuario"
		}
		return m, nil
	},
}
