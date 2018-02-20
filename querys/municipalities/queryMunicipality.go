package municipalities

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryMunicipality Query para consultar un municipio.
var QueryMunicipality = &graphql.Field{
	Type: graphqlTypes.MunicipalityType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
