package countries

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryCountries Query GraphQL para consultar la lista de paises
var QueryCountries = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.CountryType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListCountries []student.Country

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadCountries {
			db.Find(&ListCountries)

			for i := 0; i < len(ListCountries); i++ {
				ListCountries[i].Message = "#QueryCountries# : Consulta exitosa."
				ListCountries[i].Code = http.StatusAccepted
			}

		} else {
			Country := student.Country{
				Message: "#QueryCountries# : No tienes permisos para leer paises.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListCountries = append(ListCountries, Country)
		}
		return ListCountries, nil
	},
}
