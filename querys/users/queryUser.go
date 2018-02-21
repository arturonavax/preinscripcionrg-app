package users

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryUser Query para consultar un usuario.
var QueryUser = &graphql.Field{
	Type: graphqlTypes.UserType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		userFound := &models.User{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadMentions {
				err := db.Where("id = ?", idQuery).First(&userFound).Error

				if err != nil {
					userFound.Message = "#QueryUser# : Ocurrio un error."
					userFound.Code = http.StatusInternalServerError
				} else {
					userFound.Message = "#QueryUser# : Consulta exitosa."
					userFound.Code = http.StatusAccepted
				}
			} else {
				userFound.Message = "#QueryUser# : No tienes permisos para leer usuarios."
				userFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return userFound, nil
	},
}
