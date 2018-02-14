package graphqlTypes

import "github.com/graphql-go/graphql"

//TypeOfRoadType Objeto GraphQL para los tipos de via.
var TypeOfRoadType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Parish",
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
