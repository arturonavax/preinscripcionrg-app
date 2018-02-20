package teachers

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryTeacher Query para consultar un profesor.
var QueryTeacher = &graphql.Field{
	Type: graphqlTypes.TeacherType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
