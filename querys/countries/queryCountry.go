package countries

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryCountry Query para consultar un pais.
var QueryCountry = &graphql.Field{
	Type: graphqlTypes.CountryType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		countryFound := &student.Country{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadCountries {
				err := db.Where("id = ?", idQuery).First(&countryFound).Error

				if err != nil {
					countryFound.Message = "#QueryCountry# : Ocurrio un error."
					countryFound.Code = http.StatusInternalServerError
				} else {
					countryFound.Message = "#QueryCountry# : Consulta exitosa."
					countryFound.Code = http.StatusAccepted
				}
			} else {
				countryFound.Message = "#QueryCountry# : No tienes permisos para leer paises."
				countryFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return countryFound, nil
	},
}
