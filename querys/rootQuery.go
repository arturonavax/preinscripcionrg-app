package querys

import (
	"github.com/arthurnavah/PreInscripcionRG/querys/me"
	"github.com/graphql-go/graphql"
)

//RootQuery administrador de Querys
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": me.QueryMe,
		},
	},
)
