package teachers

import (
	"net/http"

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

			for i := 0; i < len(ListTeachers); i++ {
				ListTeachers[i].Message = "#QueryTeachers# : Consulta exitosa."
				ListTeachers[i].Code = http.StatusAccepted
			}
		} else {

			Teacher := student.Teacher{
				Message: "#QueryTeachers# : No tienes permisos para leer profesores.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListTeachers = append(ListTeachers, Teacher)
		}
		return ListTeachers, nil
	},
}
