package utils

import (
	"flag"
	"os"
	"strings"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/databases"
)

//CreateTables Decide si se desea Crear las tablas de la Base de Datos.
var CreateTables string

//DropTables Decide si se desea borrar las tablas de la Base de Datos.
var DropTables string

//RecreateTables Decide si se desea volver a Crear las tablas de la Base de Datos.
var RecreateTables string

//Port Contiene el Puerto en el que desea ejecutar el Servidor.
var Port int

//InitProduction Permite crear las carpetas y archivos necesarios del sistema.
var InitProduc string

//FlagControl Controla las banderas del ejecutable.
func FlagControl() {
	flag.StringVar(&CreateTables, "createTables", "NO", "Genera las tablas necesarias")
	flag.StringVar(&DropTables, "dropTables", "NO", "Borra las tablas de la base de datos")
	flag.StringVar(&RecreateTables, "recreateTables", "NO", "Recrea las tablas necesarias")
	flag.StringVar(&InitProduc, "initProduction", "NO", "Crea todas las carpetas y archivos necesarios para ejecutar la aplicacion")
	flag.IntVar(&Port, "port", 8080, "Declara el puerto para el servidor web")
	flag.Parse()

	CreateTables = strings.ToUpper(CreateTables)
	DropTables = strings.ToUpper(DropTables)
	RecreateTables = strings.ToUpper(RecreateTables)
	InitProduc = strings.ToUpper(InitProduc)

	if CreateTables == "YES" {
		databases.CreateTables()
		os.Exit(0)

	}
	if DropTables == "YES" {
		databases.DropTables()
		os.Exit(0)

	}
	if RecreateTables == "YES" {
		databases.RecreateTables()
		os.Exit(0)

	}

	if InitProduc == "YES" {
		InitProduction()
		os.Exit(0)

	}

}
