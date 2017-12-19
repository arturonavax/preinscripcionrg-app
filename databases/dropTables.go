package databases

import (
	"log"

	"gitlab.com/arthurnavah/Production-Arthur/Inventario-API/models"
)

//DropTables elimina las tablas de la base de datos.
func DropTables() {
	log.Println("---------------------------------------")
	log.Println("-         Eliminando tablas..		   -")

	//Se obtiene la conexion a la DB.
	db := GetConnectionDB()
	defer db.Close()

	//Se eliminan las tablas.
	db.DropTable(&models.People{})
	db.DropTable(&models.UserKind{})
	db.DropTable(&models.UserStatus{})
	db.DropTable(&models.User{})
	//------------------------

	log.Println("-                 -                    -")
	log.Println("- ...Finalizo la eliminacion de Tablas -")
	log.Println("----------------------------------------")
}
