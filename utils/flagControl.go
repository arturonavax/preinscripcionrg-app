package utils

import (
	"flag"
)

//RecreateTables Decide si se desea volver a Crear las tablas de la Base de Datos.
var RecreateTables string

//CreateTables Decide si se desea Crear las tablas de la Base de Datos.
var CreateTables string

//Port Contiene el Puerto en el que desea ejecutar el Servidor.
var Port int

//FlagControl Controla las banderas del ejecutable.
func FlagControl() {
	flag.StringVar(&CreateTables, "createTables", "no", "Genera las tablas necesarias")
	flag.StringVar(&RecreateTables, "recreateTables", "no", "Recrea las tablas necesarias")
	flag.IntVar(&Port, "port", 8080, "Declara el puerto para el servidor web")
	flag.Parse()
	if CreateTables == "yes" {
		createTables()
	}
	if RecreateTables == "yes" {
		recreateTables()
	}
}
