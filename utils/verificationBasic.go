package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
)

//VerificationBasic Hace las verificaciones basicas para el funcionamiento del sistema.
func VerificationBasic() {
	var permit uint8
	_, err := ioutil.ReadFile("./security/keys/private.rsa")
	if err != nil {
		log.Println("=========================================================================")
		log.Println("!! No se pudo leer el archivo privado en './security/keys/private.rsa' !!")
		log.Println("=========================================================================")
		permit++
	}

	_, err = ioutil.ReadFile("./security/keys/public.rsa")
	if err != nil {
		log.Println("========================================================================")
		log.Println("!! No se pudo leer el archivo publico en './security/keys/public.rsa' !!")
		log.Println("========================================================================")
		permit++
	}

	_, err = ioutil.ReadFile("./databases/configDB.json")
	if err != nil {
		log.Println("=================================================================================================")
		log.Println("!! No se pudo leer el archivo de configuracion de la base de datos './databases/configDB.json' !!")
		log.Println("=================================================================================================")
		permit++
	}

	if permit != 0 {
		InitProduction()
		log.Println("================================================")
		log.Println("# --> SE HAN CREADO LOS ARCHIVOS FALTANTES <-- #")
		log.Println("================================================")
		os.Exit(0)
	}
	//Comprobacion de Tablas.
	tablesExist := TablesIsExist(
		models.People{},
		models.UserStatus{},
		models.UserKind{},
		models.User{},
		student.Country{},
		student.State{},
		student.Municipality{},
		student.Institution{},
		student.Representative{},
		student.Teacher{},
		student.Parish{},
		student.Sector{},
		student.TypeOfRoad{},
		student.Father{},
		student.Mother{},
		student.Mention{},
		student.Section{},
		student.StudentCondition{},
		student.Student{},
		student.Asignature{},
	)
	//-----------------------

	if tablesExist != true {
		var t string
		log.Println("========================================================================================================")
		log.Println("!! Se detecto la ausencia de algunas Tablas necesarias en la Base de Datos, ¿Desea crearlas? (Yes/No) !!")
		fmt.Print("                    + >")
		fmt.Scan(&t)
		log.Println("========================================================================================================")
		t = strings.ToUpper(t)
		if t == "YES" {
			databases.CreateTables()

		} else {
			log.Println("=======================================================")
			log.Println("¡¡ No es posible continuar sin las tablas necesarias ¡¡")
			log.Println("=======================================================")
			os.Exit(0)

		}

	}

}

//TablesIsExist Comprueba la existencia de las tablas/estructuras que se le pasen como parametros, Devuelve un bool.
func TablesIsExist(modelss ...interface{}) bool {
	allExist := true

	//Se obtiene la conexion a la DB.
	db := databases.GetConnectionDB()
	defer db.Close()

	//Sobreescribe el slice de parametros con un booleano si existe o no.
	for i, v := range modelss {
		existTable := db.HasTable(v)

		modelss[i] = existTable

	}

	//Se recorre el slice sobreescrito, Si se encuentra un valor false significa que una tabla no existe.
	for _, v := range modelss {
		if v == false {
			allExist = false

		}

	}

	//Se devuelve la variable allExist, Si dentro de el slice de parametros habia una tabla que no existia, allExist
	// es false.
	return allExist

}
