package states

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStates Query GraphQL para consultar lista de estados
var QueryStates = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.StateType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListStates []student.State

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadStates {
			db.Find(&ListStates)

			for i := 0; i < len(ListStates); i++ {
				ListStates[i].Message = "#QueryStates# : Consulta exitosa."
				ListStates[i].Code = http.StatusAccepted
			}
		} else {
			State := student.State{
				Message: "#QueryStates# : No tienes permisos para leer estados.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListStates = append(ListStates, State)

		}
		return ListStates, nil
	},
}
