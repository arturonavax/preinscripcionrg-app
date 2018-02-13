package students

import (
	"log"
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//StudentCreate Mutation para crear Estudiantes.
var StudentCreate = &graphql.Field{
	Type: graphqlTypes.StudentType,
	Args: graphql.FieldConfigArgument{
		"countryOfBirthID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"stateOfBirthID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"municOfBirthID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"municipalityID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"institutionOfOriginID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"parishID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"sectorID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"typeOfRoadID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"mentionID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"sectionID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"studentConditionID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"representativeID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"teacherID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"SIGECOD": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"firstName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"lastName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"ciType": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"ci": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"motherFirstName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"motherLastName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"motherEmail": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"motherPhoneNumber": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"motherCIType": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"motherCI": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"fatherFirstName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fatherLastName": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fatherEmail": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fatherPhoneNumber": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fatherCIType": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"fatherCI": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"healthProblem": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"healthProblemE": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"address": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"scholarship": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"canaima": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"conditionOfHousing": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"year": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"age": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"weight": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"repeatAsignature": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"pendingAsignature": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"regular": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"inscriptionDate": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		newMother := &student.Mother{}
		newFather := &student.Father{}
		newStudent := &student.Student{}
		m := &models.CreateMessage{}
		kind := &models.UserKind{}

		//Se obtiene el ID del usuario enviado en el Contexto.
		userID := p.Context.Value("user").(models.User).ID

		//Se obtiene el ID del tipo de usuario enviado en el contexto.
		kindID := p.Context.Value("user").(models.User).KindID

		//Se pobla newStudent con la informacion enviada como argumentos.
		newStudent.UserID = userID
		newStudent.CountryOfBirthID, _ = p.Args["countryOfBirthID"].(int)
		newStudent.StateOfBirthID, _ = p.Args["stateOfBirthID"].(int)
		newStudent.MunicOfBirthID, _ = p.Args["municOfBirthID"].(int)
		newStudent.MunicipalityID, _ = p.Args["municipalityID"].(int)
		newStudent.InstitutionOfOriginID, _ = p.Args["institutionOfOriginID"].(int)
		newStudent.ParishID, _ = p.Args["parishID"].(int)
		newStudent.SectorID, _ = p.Args["sectorID"].(int)
		newStudent.TypeOfRoadID, _ = p.Args["typeOfRoadID"].(int)
		newStudent.MentionID, _ = p.Args["mentionID"].(int)
		newStudent.SectionID, _ = p.Args["sectionID"].(int)
		newStudent.StudentConditionID, _ = p.Args["studentConditionID"].(int)
		newStudent.RepresentativeID, _ = p.Args["representativeID"].(int)
		newStudent.TeacherID, _ = p.Args["teacherID"].(int)

		newStudent.SIGECOD, _ = p.Args["SIGECOD"].(string)

		newStudent.MotherFirstName, _ = p.Args["motherFirstName"].(string)
		newStudent.MotherLastName, _ = p.Args["motherLastName"].(string)
		newStudent.MotherEmail, _ = p.Args["motherEmail"].(string)
		newStudent.MotherPhoneNumber, _ = p.Args["motherPhoneNumber"].(string)
		newStudent.MotherCIType, _ = p.Args["motherCIType"].(string)
		newStudent.MotherCI, _ = p.Args["motherCI"].(int)

		newStudent.FatherFirstName, _ = p.Args["fatherFirstName"].(string)
		newStudent.FatherLastName, _ = p.Args["fatherLastName"].(string)
		newStudent.FatherEmail, _ = p.Args["fatherEmail"].(string)
		newStudent.FatherPhoneNumber, _ = p.Args["fatherPhoneNumber"].(string)
		newStudent.FatherCIType, _ = p.Args["fatherCIType"].(string)
		newStudent.FatherCI, _ = p.Args["fatherCI"].(int)

		newStudent.FirstName, _ = p.Args["firstName"].(string)

		newStudent.LastName, _ = p.Args["lastName"].(string)

		newStudent.CIType, _ = p.Args["ciType"].(string)

		newStudent.CI, _ = p.Args["ci"].(int)

		newStudent.DateOfBirth, _ = p.Args["dateOfBirth"].(string)

		newStudent.Gender, _ = p.Args["gender"].(string)

		newStudent.HealthProblem, _ = p.Args["healthProblem"].(bool)

		newStudent.HealthProblemE, _ = p.Args["healthProblemE"].(string)

		newStudent.Email, _ = p.Args["email"].(string)

		newStudent.PhoneNumber, _ = p.Args["phoneNumber"].(string)

		newStudent.Address, _ = p.Args["address"].(string)

		newStudent.Scholarship, _ = p.Args["scholarship"].(bool)

		newStudent.Canaima, _ = p.Args["canaima"].(bool)

		newStudent.ConditionOfHousing, _ = p.Args["conditionOfHousing"].(string)

		newStudent.Year, _ = p.Args["year"].(int)

		newStudent.Age, _ = p.Args["age"].(int)

		newStudent.Size, _ = p.Args["size"].(string)

		newStudent.Weight, _ = p.Args["weight"].(int)

		newStudent.RepeatAsignature, _ = p.Args["repeatAsignature"].(string)

		newStudent.PendingAsignature, _ = p.Args["pendingAsignature"].(string)

		newStudent.Regular, _ = p.Args["regular"].(bool)

		newStudent.InscriptionDate, _ = p.Args["inscriptionDate"].(string)

		newMother.FirstName = newStudent.MotherFirstName
		newMother.LastName = newStudent.MotherLastName
		newMother.Email = newStudent.MotherEmail
		newMother.PhoneNumber = newStudent.MotherPhoneNumber
		newMother.CIType = newStudent.MotherCIType
		newMother.CI = newStudent.MotherCI

		newFather.FirstName = newStudent.FatherFirstName
		newFather.LastName = newStudent.FatherLastName
		newFather.Email = newStudent.FatherEmail
		newFather.PhoneNumber = newStudent.FatherPhoneNumber
		newFather.CIType = newStudent.FatherCIType
		newFather.CI = newStudent.FatherCI

		//Se instancia una conexion a la base de datos.
		db := databases.GetConnectionDB()
		defer db.Close()

		//Se realiza un Query para obtener el tipo de usuario.
		db.Where("id = ?", kindID).First(&kind)

		//Permisos para crear estudiantes.
		if kind.CreateStudents {
			tx := db.Begin()

			err := tx.Create(&newMother).Error
			if err != nil {
				m.Message = "#StudentCreate# : Ocurrio un Error al registrar a la madre."
				m.Code = http.StatusBadGateway
				tx.Rollback()
			}

			err = tx.Create(&newFather).Error
			if err != nil {
				m.Message = "#StudentCreate# : Ocurrio un Error al registrar al padre."
				m.Code = http.StatusBadGateway
				tx.Rollback()
			}

			newStudent.MotherID = newMother.ID
			newStudent.FatherID = newFather.ID

			//Se crea el Estudiante en la base de datos.
			err = tx.Create(&newStudent).Error
			if err != nil {
				m.Message = "#StudentCreate# : Ocurrio un Error al registrar el estudiante."
				m.Code = http.StatusBadGateway
			} else {
				tx.Commit()
				log.Printf("$ UserID:%d { Registro a '%s %s' como nuevo estudiante (StudentID:%d) } $\n",
					userID,
					newStudent.FirstName,
					newStudent.LastName,
					newStudent.ID,
				)
				m.ID = newStudent.ID
				m.Code = http.StatusCreated
				m.Message = "#StudentCreate# : El estudiante se creo exitosamente!."
			}
		} else {
			log.Printf("$ UserID:%d { No tiene permisos para crear estudiantes } $\n", userID)
			m.Code = http.StatusNonAuthoritativeInfo
			m.Message = "#StudentCreate# : No tienes permisos para crear estudiantes."
		}

		//Devuelve el mensaje al usuario
		return m, nil
	},
}
