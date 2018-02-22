package institutions

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryInstitution Query para consultar una institucion.
var QueryInstitution = &graphql.Field{
	Type: graphqlTypes.InstitutionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		institutionFound := &student.Institution{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadInstitutions {
				err := db.Where("id = ?", idQuery).First(&institutionFound).Error

				if err != nil {
					institutionFound.Message = "#QueryInstitution# : Ocurrio un error."
					institutionFound.Code = http.StatusInternalServerError
				} else {
					institutionFound.Message = "#QueryInstitution# : Consulta exitosa."
					institutionFound.Code = http.StatusAccepted
				}
			} else {
				institutionFound.Message = "#QueryInstitution# : No tienes permisos para leer institutiones."
				institutionFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return institutionFound, nil
	},
}
