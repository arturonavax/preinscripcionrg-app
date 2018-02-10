package users

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

var QueryUsers = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.UserType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

	},
}
