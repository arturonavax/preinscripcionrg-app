package schemas

import (
	"github.com/graphql-go/graphql"
	"gitlab.com/arthurnavah/Production-Arthur/InventSoft2/mutations"
	"gitlab.com/arthurnavah/Production-Arthur/InventSoft2/querys"
)

//SchemaInit Contiene el Schema GraphQL con los Querys y Mutations.
var SchemaInit, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    querys.RootQuery,
		Mutation: mutations.RootMutation,
	},
)
