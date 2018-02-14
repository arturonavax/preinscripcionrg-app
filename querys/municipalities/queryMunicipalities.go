package municipalities

import (
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

		if kind.ReadMentions {
			db.Find(&ListMunicipalities)
		}
		return ListMunicipalities, nil
	},
}
