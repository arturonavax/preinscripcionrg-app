package databases

import (
	"log"
	"github.com/arthurnavah/PreInscripcionRG-API/models"
)

//Migrate Crea las tablas necesarias.
func createTables() {
	log.Println("---------------------------------------")
	log.Println("- Generando Tablas...                 -")

	db := GetConnectionDB()
	defer db.Close()
	db.CreateTable(&models.People{})
	db.CreateTable(&models.UserKind{})
	db.CreateTable(&models.UserStatus{})
	db.CreateTable(&models.User{})

	db.Model(&models.User{}).Related(&models.People{})

	db.Model(&models.User{}).AddForeignKey("people_id", "peoples(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.User{}).AddForeignKey("kind_id", "user_kinds(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.User{}).AddForeignKey("status_id", "user_statuses(id)", "RESTRICT", "RESTRICT")

	kind := &models.UserKind{}
	kind.Name = "ESPECTADOR"
	kind.Description = "No tiene ningun Permiso"

	status := &models.UserStatus{}
	status.Name = "ACTIVO"
	status.Description = "Usuario activo"

	db.Create(&kind)
	db.Create(&status)

	log.Println("- ...Finalizo la Generacion de Tablas -")
	log.Println("---------------------------------------")
	log.Println(" 	")
}
