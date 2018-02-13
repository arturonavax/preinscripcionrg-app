package graphqlTypes

import "github.com/graphql-go/graphql"

//SectionType Objeto GraphQL para secciones.
var SectionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Section",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"code": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
