package schemas

import (
	"github.com/arthurnavah/PreInscripcionRG-API/mutations"
	"github.com/arthurnavah/PreInscripcionRG-API/querys"
	"github.com/graphql-go/graphql"
)

//SchemaInit Contiene el Schema GraphQL con los Querys y Mutations.
var SchemaInit, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    querys.RootQuery,
		Mutation: mutations.RootMutation,
	},
)
