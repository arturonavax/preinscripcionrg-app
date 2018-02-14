package students

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStudents Query GraphQL para consultar lista de estudiantes
var QueryStudents = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.StudentType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListStudent []student.Student

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadStudents {
			db.Find(&ListStudent)
		}
		return ListStudent, nil
	},
}
