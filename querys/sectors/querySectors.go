package sectors

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QuerySectors Query GraphQL para consultar lista de Sectores
var QuerySectors = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.SectorType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListSectors []student.Sector

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadSectors {
			db.Find(&ListSectors)

			for i := 0; i < len(ListSectors); i++ {
				ListSectors[i].Message = "#QuerySectors# : Consulta exitosa."
				ListSectors[i].Code = http.StatusAccepted
			}
		} else {
			Sector := student.Sector{
				Message: "#QuerySectors# : No tienes permisos para leer sectores.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListSectors = append(ListSectors, Sector)

		}
		return ListSectors, nil
	},
}
