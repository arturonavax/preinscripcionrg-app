package graphqlTypes

import "github.com/graphql-go/graphql"

//StudentConditionType Objeto GraphQL para condiciones de estudiantes.
var StudentConditionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "StudentCondition",
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
