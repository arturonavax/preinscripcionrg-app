package graphqlTypes

import "github.com/graphql-go/graphql"

//StateType Objeto GraphQL para estados.
var StateType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Parish",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.Int,
			},
			"countryID": &graphql.Field{
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
