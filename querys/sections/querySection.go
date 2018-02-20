package sections

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
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
		return nil, nil
	},
}
