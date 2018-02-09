package me

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

var QueryMe = &graphql.Field{
	Type: graphqlTypes.UserType,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var user = models.User{}
		var people = models.People{}
		userID := p.Context.Value("user").(models.User).ID
		peopleID := p.Context.Value("user").(models.User).PeopleID

		//Se instancia una conexion a la base de datos.
		db := databases.GetConnectionDB()
		defer db.Close()

		//Se hace un Query para obtener la informacion de este Usuario.
		db.Where("id = ?", userID).First(&user)
		db.Where("id = ?", peopleID).First(&people)

		user.FirstName = people.FirstName
		user.Address = people.Address
		user.Email = people.Email
		user.LastName = people.LastName
		user.PhoneNumber = people.PhoneNumber

		user.Password = ""
		//Devolvemos el usuario
		return user, nil
	},
}
