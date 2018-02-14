package typeOfRoads

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryTypeOfRoads Query GraphQL para consultar lista de tipos de vias
var QueryTypeOfRoads = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.TypeOfRoadType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListTypeOfRoads []student.TypeOfRoad

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadTypeOfRoads {
			db.Find(&ListTypeOfRoads)
		}
		return ListTypeOfRoads, nil
	},
}
