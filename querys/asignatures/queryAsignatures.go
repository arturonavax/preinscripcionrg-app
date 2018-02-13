package asignatures

import (
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
		}
		return ListAsignatures, nil
	},
}
