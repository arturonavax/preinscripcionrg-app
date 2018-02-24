package conditionOfHousing

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryConditionOfHousings Query GraphQL para consultar lista de condiciones de vivienda
var QueryConditionOfHousings = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.StudentConditionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListConditionOfHousing []student.ConditionOfHousing

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadStudentConditions {
			db.Find(&ListConditionOfHousing)

			for i := 0; i < len(ListConditionOfHousing); i++ {
				ListConditionOfHousing[i].Message = "#QueryConditionOfHousings# : Consulta exitosa."
				ListConditionOfHousing[i].Code = http.StatusAccepted
			}
		} else {

			ConditionOfHousing := student.ConditionOfHousing{
				Message: "#QueryConditionOfHousings# : No tienes permisos para leer condiciones de viviendas.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListConditionOfHousing = append(ListConditionOfHousing, ConditionOfHousing)
		}
		return ListConditionOfHousing, nil
	},
}
