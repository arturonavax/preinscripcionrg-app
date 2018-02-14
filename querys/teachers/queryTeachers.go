package teachers

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryTeachers Query GraphQL para consultar lista de profesores
var QueryTeachers = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.TeacherType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListTeachers []student.Teacher

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadTeachers {
			db.Find(&ListTeachers)
		}
		return ListTeachers, nil
	},
}
