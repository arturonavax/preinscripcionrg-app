package municipalities

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryMunicipality Query para consultar un municipio.
var QueryMunicipality = &graphql.Field{
	Type: graphqlTypes.MunicipalityType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		municipalityFound := &student.Municipality{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadMunicipalities {
				err := db.Where("id = ?", idQuery).First(&municipalityFound).Error

				if err != nil {
					municipalityFound.Message = "#QueryMunicipality# : Ocurrio un error."
					municipalityFound.Code = http.StatusInternalServerError
				} else {
					municipalityFound.Message = "#QueryMunicipality# : Consulta exitosa."
					municipalityFound.Code = http.StatusAccepted
				}
			} else {
				municipalityFound.Message = "#QueryMunicipality# : No tienes permisos para leer municipios."
				municipalityFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return municipalityFound, nil
	},
}
