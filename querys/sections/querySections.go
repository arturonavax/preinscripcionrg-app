package sections

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QuerySections Query GraphQL para consultar lista de secciones
var QuerySections = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.SectionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListSections []student.Section

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadSections {
			db.Find(&ListSections)

			for i := 0; i < len(ListSections); i++ {
				ListSections[i].Message = "#QuerySections# : Consulta exitosa."
				ListSections[i].Code = http.StatusAccepted
			}
		} else {
			Section := student.Section{
				Message: "#QuerySections# : No tienes permisos para leer secciones.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListSections = append(ListSections, Section)
		}
		return ListSections, nil
	},
}
