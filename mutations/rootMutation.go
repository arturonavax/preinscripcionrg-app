package mutations

import (
	"github.com/graphql-go/graphql"
)

//RootMutation administrador de mutations.
var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: graphql.Fields{},
	},
)
