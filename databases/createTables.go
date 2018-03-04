package databases

import (
	"log"

	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
)

//CreateTables Crea las tablas necesarias.
func CreateTables() {
	log.Println("=======================================")
	log.Println("|        Generando Tablas...          |")

	//Se obtiene la conexion a la DB.
	db := GetConnectionDB()
	defer db.Close()

	//Se crean las tablas.
	db.CreateTable(&models.People{})
	db.CreateTable(&models.UserStatus{})
	db.CreateTable(&models.UserKind{})
	db.CreateTable(&models.User{})
	db.CreateTable(&student.Country{})
	db.CreateTable(&student.State{})
	db.CreateTable(&student.Municipality{})
	db.CreateTable(&student.Institution{})
	db.CreateTable(&student.Representative{})
	db.CreateTable(&student.Teacher{})
	db.CreateTable(&student.Parish{})
	db.CreateTable(&student.Sector{})
	db.CreateTable(&student.TypeOfRoad{})
	db.CreateTable(&student.Father{})
	db.CreateTable(&student.Mother{})
	db.CreateTable(&student.Mention{})
	db.CreateTable(&student.Section{})
	db.CreateTable(&student.ConditionOfHousing{})
	db.CreateTable(&student.Student{})
	db.CreateTable(&student.StudentStatus{})
	db.CreateTable(&student.Asignature{})
	//----------------------

	//Relaciones y llaves foraneas.
	db.Model(&models.User{}).Related(&models.People{})

	db.Model(&models.User{}).AddForeignKey("people_id", "peoples(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.User{}).AddForeignKey("kind_id", "user_kinds(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.User{}).AddForeignKey("status_id", "user_statuses(id)", "RESTRICT", "RESTRICT")
	//------------------------------

	//Records basicos necesarios.
	kind := &models.UserKind{}
	kind.Name = "REPRESENTANTE"
	kind.Description = "Tiene los permisos basicos para poder utilizar la interfaz superficialmente y registrar usuarios"

	status := &models.UserStatus{}
	status.Name = "ACTIVO"
	status.Description = "Usuario activo"

	studentStatus := &student.StudentStatus{}
	studentStatus.Name = "EN PROCESO"
	studentStatus.Description = "Estudiantes en proceso de ser aceptados o rechazados"

	studentStatus2 := &student.StudentStatus{}
	studentStatus2.Name = "ACEPTADO"
	studentStatus2.Description = "Estudiantes aceptados para ser inscritos"

	studentStatus3 := &student.StudentStatus{}
	studentStatus3.Name = "RECHAZADO"
	studentStatus3.Description = "Estudiantes rechazados para ser inscritos"

	db.Create(&kind)
	db.Create(&status)
	db.Create(&studentStatus)
	db.Create(&studentStatus2)
	db.Create(&studentStatus3)
	//---------------------------

	log.Println("|                 ...                 |")
	log.Println("| ...Finalizo la Generacion de Tablas |")
	log.Println("=======================================")
}
