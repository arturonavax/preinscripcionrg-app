package studentConditions

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStudentCondition Query para consultar una condicion de estudiante.
var QueryStudentCondition = &graphql.Field{
	Type: graphqlTypes.StudentConditionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		studentConditionFound := &student.StudentCondition{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStudentConditions {
				err := db.Where("id = ?", idQuery).First(&studentConditionFound).Error

				if err != nil {
					studentConditionFound.Message = "#QueryStudentCondition# : Ocurrio un error."
					studentConditionFound.Code = http.StatusInternalServerError
				} else {
					studentConditionFound.Message = "#QueryStudentCondition# : Consulta exitosa."
					studentConditionFound.Code = http.StatusAccepted
				}
			} else {
				studentConditionFound.Message = "#QueryStudentCondition# : No tienes permisos para leer condiciones de estudiantes."
				studentConditionFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return studentConditionFound, nil
	},
}
