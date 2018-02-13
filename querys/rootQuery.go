package querys

import (
	"github.com/arthurnavah/PreInscripcionRG/querys/asignatures"
	"github.com/arthurnavah/PreInscripcionRG/querys/me"
	"github.com/arthurnavah/PreInscripcionRG/querys/students"
	"github.com/arthurnavah/PreInscripcionRG/querys/users"
	"github.com/graphql-go/graphql"
)

//RootQuery administrador de Querys
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me":          me.QueryMe,
			"users":       users.QueryUsers,
			"students":    students.QueryStudents,
			"student":     students.QueryStudent,
			"asignatures": asignatures.QueryAsignatures,
		},
	},
)
