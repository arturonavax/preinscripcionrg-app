package mutations

import (
	"github.com/arthurnavah/PreInscripcionRG/mutations/me"
	"github.com/arthurnavah/PreInscripcionRG/mutations/students"
	"github.com/graphql-go/graphql"
)

//RootMutation administrador de mutations.
var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//Mutations de Estudiantes.
			"studentC": students.StudentCreate,
			//Mutations del Usuario.
			"meU": me.MeUpdate,
		},
	},
)
