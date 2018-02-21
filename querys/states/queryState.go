package states

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryState Query para consultar un estado.
var QueryState = &graphql.Field{
	Type: graphqlTypes.StateType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		stateFound := &student.State{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStates {
				err := db.Where("id = ?", idQuery).First(&stateFound).Error

				if err != nil {
					stateFound.Message = "#QueryState# : Ocurrio un error."
					stateFound.Code = http.StatusInternalServerError
				} else {
					stateFound.Message = "#QueryState# : Consulta exitosa."
					stateFound.Code = http.StatusAccepted
				}
			} else {
				stateFound.Message = "#QueryState# : No tienes permisos para leer estados."
				stateFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return stateFound, nil
	},
}
