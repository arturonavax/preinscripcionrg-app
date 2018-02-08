package schemas

import (
	"github.com/arthurnavah/PreInscripcionRG/mutations"
	"github.com/arthurnavah/PreInscripcionRG/querys"
	"github.com/graphql-go/graphql"
)

//SchemaInit Contiene el Schema GraphQL con los Querys y Mutations.
var SchemaInit, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    querys.RootQuery,
		Mutation: mutations.RootMutation,
	},
)
