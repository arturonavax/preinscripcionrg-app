package users

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryUsers Query GraphQL para consultar lista de Usuarios
var QueryUsers = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.UserType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListUsers []models.User

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadUsers {
			db.Find(&ListUsers)
		}

		return ListUsers, nil
	},
}
