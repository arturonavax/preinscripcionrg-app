package studentConditions

import (
	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryStudentConditions Query GraphQL para consultar lista de condiciones de estudiantes
var QueryStudentConditions = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.StudentConditionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListStudentConditions []student.StudentCondition

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadStudentConditions {
			db.Find(&ListStudentConditions)
		}
		return ListStudentConditions, nil
	},
}
