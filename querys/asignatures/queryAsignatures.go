package asignatures

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryAsignatures Query GraphQL para consultar la lista de asignaturas
var QueryAsignatures = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.AsignatureType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListAsignatures []student.Asignature

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadAsignatures {
			db.Find(&ListAsignatures)

			for i := 0; i < len(ListAsignatures); i++ {
				ListAsignatures[i].Message = "#QueryAsignatures# : Consulta exitosa."
				ListAsignatures[i].Code = http.StatusAccepted
			}

		} else {
			Asignature := student.Asignature{
				Message: "#QueryAsignatures# : No tienes permisos para leer asignaturas.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListAsignatures = append(ListAsignatures, Asignature)
		}
		return ListAsignatures, nil
	},
}
