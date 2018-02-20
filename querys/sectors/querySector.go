package sectors

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QuerySector Query para consultar un sector.
var QuerySector = &graphql.Field{
	Type: graphqlTypes.SectorType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
