package users

import (
	"net/http"

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

			for i := 0; i < len(ListUsers); i++ {
				ListUsers[i].Message = "#QueryUsers# : Consulta exitosa."
				ListUsers[i].Code = http.StatusAccepted
			}
		} else {
			User := models.User{
				Message: "#QueryUsers# : No tienes permisos para leer usuarios.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListUsers = append(ListUsers, User)

		}

		return ListUsers, nil
	},
}
