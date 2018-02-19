package municipalities

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryMunicipalities Query GraphQL para consultar lista de municipios
var QueryMunicipalities = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.MunicipalityType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListMunicipalities []student.Municipality

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadMunicipalities {
			db.Find(&ListMunicipalities)

			for i := 0; i < len(ListMunicipalities); i++ {
				ListMunicipalities[i].Message = "#QueryMunicipalities# : Consulta exitosa."
				ListMunicipalities[i].Code = http.StatusAccepted
			}
		} else {

			Municipality := student.Municipality{
				Message: "#QueryMunicipalities# : No tienes permisos para leer municipios.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListMunicipalities = append(ListMunicipalities, Municipality)
		}
		return ListMunicipalities, nil
	},
}
