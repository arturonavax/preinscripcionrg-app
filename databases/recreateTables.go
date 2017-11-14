package databases

import (
	"log"

	"github.com/arthurnavah/PreInscripcionRG-API/models"
)

//recreateTables Elimina y Crea las tablas de la base de datos.
func recreateTables() {
	log.Println("---------------------------------------")
	log.Println("- Recreando tablas...                 -")

	db := GetConnectionDB()
	defer db.Close()
	//Se borran las tablas.
	db.DropTable(&models.User{})
	db.DropTable(&models.People{})
	db.DropTable(&models.UserKind{})
	db.DropTable(&models.UserStatus{})

	createTables()
}
