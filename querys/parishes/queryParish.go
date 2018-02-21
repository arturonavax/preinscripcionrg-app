package parishes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryParish Query para consultar una parroquia.
var QueryParish = &graphql.Field{
	Type: graphqlTypes.MunicipalityType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		parishFound := &student.Parish{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadParishes {
				err := db.Where("id = ?", idQuery).First(&parishFound).Error

				if err != nil {
					parishFound.Message = "#QueryParish# : Ocurrio un error."
					parishFound.Code = http.StatusInternalServerError
				} else {
					parishFound.Message = "#QueryParish# : Consulta exitosa."
					parishFound.Code = http.StatusAccepted
				}
			} else {
				parishFound.Message = "#QueryParish# : No tienes permisos para leer parroquias."
				parishFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return parishFound, nil
	},
}
