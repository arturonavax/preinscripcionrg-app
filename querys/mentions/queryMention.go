package mentions

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryMention Query para consultar una mencion.
var QueryMention = &graphql.Field{
	Type: graphqlTypes.MentionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
