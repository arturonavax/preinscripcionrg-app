package sectors

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QuerySector Query para consultar un sector.
var QuerySector = &graphql.Field{
	Type: graphqlTypes.SectorType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		sectorFound := &student.Sector{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadSectors {
				err := db.Where("id = ?", idQuery).First(&sectorFound).Error

				if err != nil {
					sectorFound.Message = "#QuerySector# : Ocurrio un error."
					sectorFound.Code = http.StatusInternalServerError
				} else {
					sectorFound.Message = "#QuerySector# : Consulta exitosa."
					sectorFound.Code = http.StatusAccepted
				}
			} else {
				sectorFound.Message = "#QuerySector# : No tienes permisos para leer sectores."
				sectorFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return sectorFound, nil
	},
}
