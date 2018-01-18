package students

import (
	"github.com/graphql-go/graphql"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models/graphqlTypes"
)

//StudentCreate Mutation para crear Estudiantes.
var StudentCreate = &graphql.Field{
	Type: graphqlTypes.StudentType,
	Args: graphql.FieldConfigArgument{},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

	},
}
