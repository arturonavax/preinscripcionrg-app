package typeOfRoads

import (
	"net/http"

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

			for i := 0; i < len(ListTypeOfRoads); i++ {
				ListTypeOfRoads[i].Message = "#QueryTypeOfRoads# : Consulta exitosa."
				ListTypeOfRoads[i].Code = http.StatusAccepted
			}
		} else {
			TypeOfRoad := student.TypeOfRoad{
				Message: "#QueryTypeOfRoads# : No tienes permisos para leer tipos de vias.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListTypeOfRoads = append(ListTypeOfRoads, TypeOfRoad)

		}
		return ListTypeOfRoads, nil
	},
}
