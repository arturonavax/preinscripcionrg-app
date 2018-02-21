package teachers

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
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
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		teacherFound := &student.Teacher{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadTeachers {
				err := db.Where("id = ?", idQuery).First(&teacherFound).Error

				if err != nil {
					teacherFound.Message = "#QueryTeacher# : Ocurrio un error."
					teacherFound.Code = http.StatusInternalServerError
				} else {
					teacherFound.Message = "#QueryTeacher# : Consulta exitosa."
					teacherFound.Code = http.StatusAccepted
				}
			} else {
				teacherFound.Message = "#QueryTeacher# : No tienes permisos para leer profesores."
				teacherFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return teacherFound, nil
	},
}
