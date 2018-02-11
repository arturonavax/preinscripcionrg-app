/*package users

import (
	"github.com/graphql-go/graphql"
)

var UserUpdate = &graphql.Field{
	Type: graphqlTypes.UserType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"phoneNumber": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		m := &models.Message{}
		userFound := &models.User{}
	}
}*/