package typeOfRoads

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryTypeOfRoad Query para consultar un tipo de via.
var QueryTypeOfRoad = &graphql.Field{
	Type: graphqlTypes.TypeOfRoadType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		typeOfRoadFound := &student.TypeOfRoad{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadTypeOfRoads {
				err := db.Where("id = ?", idQuery).First(&typeOfRoadFound).Error

				if err != nil {
					typeOfRoadFound.Message = "#QueryTypeOfRoad# : Ocurrio un error."
					typeOfRoadFound.Code = http.StatusInternalServerError
				} else {
					typeOfRoadFound.Message = "#QueryTypeOfRoad# : Consulta exitosa."
					typeOfRoadFound.Code = http.StatusAccepted
				}
			} else {
				typeOfRoadFound.Message = "#QueryTypeOfRoad# : No tienes permisos para leer tipos de vias."
				typeOfRoadFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return typeOfRoadFound, nil
	},
}
