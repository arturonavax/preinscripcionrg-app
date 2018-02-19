package asignatures

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryAsignature Query para consultar una asignatura.
var QueryAsignature = &graphql.Field{
	Type: graphqlTypes.AsignatureType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		asignatureFound := &student.Asignature{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadAsignatures {
				err := db.Where("id = ?", idQuery).First(&asignatureFound).Error

				if err != nil {
					asignatureFound.Message = "#QueryAsignature# : Ocurrio un error."
					asignatureFound.Code = http.StatusInternalServerError
				} else {
					asignatureFound.Message = "#QueryAsignature# : Consulta exitosa."
					asignatureFound.Code = http.StatusAccepted
				}
			} else {
				asignatureFound.Message = "#QueryAsignature# : No tienes permisos para leer asignaturas."
				asignatureFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return asignatureFound, nil
	},
}
