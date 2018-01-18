package mutations

import (
	"github.com/graphql-go/graphql"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/mutations/students"
)

//RootMutation administrador de mutations.
var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//Mutations de Estudiantes.
			"studentC": students.StudentCreate,
		},
	},
)
