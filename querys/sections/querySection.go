package sections

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QuerySection Query para consultar una seccion.
var QuerySection = &graphql.Field{
	Type: graphqlTypes.SectionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		sectionFound := &student.Section{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadSections {
				err := db.Where("id = ?", idQuery).First(&sectionFound).Error

				if err != nil {
					sectionFound.Message = "#QuerySection# : Ocurrio un error."
					sectionFound.Code = http.StatusInternalServerError
				} else {
					sectionFound.Message = "#QuerySection# : Consulta exitosa."
					sectionFound.Code = http.StatusAccepted
				}
			} else {
				sectionFound.Message = "#QuerySection# : No tienes permisos para leer secciones."
				sectionFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return sectionFound, nil
	},
}
