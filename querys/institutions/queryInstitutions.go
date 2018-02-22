package institutions

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryInstitutions Query GraphQL para consultar lista de instituciones
var QueryInstitutions = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.InstitutionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListInstitutions []student.Institution

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadInstitutions {
			db.Find(&ListInstitutions)

			for i := 0; i < len(ListInstitutions); i++ {
				ListInstitutions[i].Message = "#QueryInstitutions# : Consulta exitosa."
				ListInstitutions[i].Code = http.StatusAccepted
			}

		} else {
			Institution := student.Institution{
				Message: "#QueryInstitutions# : No tienes permisos para leer instituciones.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListInstitutions = append(ListInstitutions, Institution)

		}
		return ListInstitutions, nil
	},
}
