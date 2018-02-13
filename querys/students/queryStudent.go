package students

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

var QueryStudent = &graphql.Field{
	Type: graphqlTypes.StudentType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		studentFound := &student.Student{}
		motherFound := &student.Mother{}
		fatherFound := &student.Father{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadStudents {
				err := db.Where("id = ?", idQuery).First(&studentFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando el estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				motherFound.ID = studentFound.MotherID
				fatherFound.ID = studentFound.FatherID

				err = db.Where("id = ?", motherFound.ID).First(&motherFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando la madre del estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				err = db.Where("id = ?", fatherFound.ID).First(&fatherFound).Error
				if err != nil {
					studentFound.Message = "#QueryStudent# : Ocurrio un error buscando el padre del estudiante."
					studentFound.Code = http.StatusBadGateway
				}

				studentFound.MotherFirstName = motherFound.FirstName
				studentFound.MotherLastName = motherFound.LastName
				studentFound.MotherEmail = motherFound.Email
				studentFound.MotherPhoneNumber = motherFound.PhoneNumber
				studentFound.MotherCIType = motherFound.CIType
				studentFound.MotherCI = motherFound.CI

				studentFound.FatherFirstName = fatherFound.FirstName
				studentFound.FatherLastName = fatherFound.LastName
				studentFound.FatherEmail = fatherFound.Email
				studentFound.FatherPhoneNumber = fatherFound.PhoneNumber
				studentFound.FatherCIType = fatherFound.CIType
				studentFound.FatherCI = fatherFound.CI

			} else {
				studentFound.Message = "#QueryStudent# : No tienes permisos para leer estudiantes."
				studentFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return studentFound, nil
	},
}
