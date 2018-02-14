package graphqlTypes

import "github.com/graphql-go/graphql"

//MentionType Objeto GraphQL para menciones.
var MentionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mention",
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
