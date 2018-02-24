package conditionOfHousing

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryConditionOfHousing Query para consultar una condicion de vivienda.
var QueryConditionOfHousing = &graphql.Field{
	Type: graphqlTypes.StudentConditionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		conditionOfHousingFound := &student.ConditionOfHousing{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStudentConditions {
				err := db.Where("id = ?", idQuery).First(&conditionOfHousingFound).Error

				if err != nil {
					conditionOfHousingFound.Message = "#QueryConditionOfHousing# : Ocurrio un error."
					conditionOfHousingFound.Code = http.StatusInternalServerError
				} else {
					conditionOfHousingFound.Message = "#QueryConditionOfHousing# : Consulta exitosa."
					conditionOfHousingFound.Code = http.StatusAccepted
				}
			} else {
				conditionOfHousingFound.Message = "#QueryConditionOfHousing# : No tienes permisos para leer condiciones de vivienda."
				conditionOfHousingFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return conditionOfHousingFound, nil
	},
}
