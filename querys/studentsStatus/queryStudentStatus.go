package studentsStatus

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStudentStatus Query para consultar un estado de estudiante.
var QueryStudentStatus = &graphql.Field{
	Type: graphqlTypes.StudentStatusType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		studentStatusFound := &student.StudentStatus{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStudentsStatus {
				err := db.Where("id = ?", idQuery).First(&studentStatusFound).Error

				if err != nil {
					studentStatusFound.Message = "#QueryStudentStatus# : Ocurrio un error."
					studentStatusFound.Code = http.StatusInternalServerError
				} else {
					studentStatusFound.Message = "#QueryStudentStatus# : Consulta exitosa."
					studentStatusFound.Code = http.StatusAccepted
				}
			} else {
				studentStatusFound.Message = "#QueryStudentStatus# : No tienes permisos para leer estados de estudiantes."
				studentStatusFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return studentStatusFound, nil
	},
}
