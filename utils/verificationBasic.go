package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/arthurnavah/PreInscripcionRG-API/configurations"
	"github.com/arthurnavah/PreInscripcionRG-API/models"
)

//VerificationBasic Hace las verificaciones basicas para el Servidor
func VerificationBasic() {
	TablesVerification(
		&models.User{},
		&models.People{},
		&models.UserStatus{},
		&models.UserKind{},
	)
}

//TablesVerification Verifica que existan las tablas necesarias.
func TablesVerification(
	model1 *models.User,
	model2 *models.People,
	model3 *models.UserStatus,
	model4 *models.UserKind,
) {

	db := configurations.GetConnectionDB()
	defer db.Close()

	existTable1 := db.HasTable(&model1)
	existTable2 := db.HasTable(&model2)
	existTable3 := db.HasTable(&model3)
	existTable4 := db.HasTable(&model4)

	if !existTable1 || !existTable2 || !existTable3 || !existTable4 {
		var decision string
		fmt.Print("Algunas de las Tablas necesarias no existen, ¿Desea Crearlas? (s/n) : ")
		fmt.Scan(&decision)

		decision = fmt.Sprintf(strings.ToUpper(decision))

		if decision == "S" {
			createTables()
		} else {
			log.Panicln("Tablas necesarias no existen")
			return
		}
	}
}
