package main

import (
	"github.com/arthurnavah/PreInscripcionRG/routes"
	"github.com/arthurnavah/PreInscripcionRG/utils"
)

func main() {
	//FlagControl Controla las banderas para el ejecutable.
	//	--createTables yes ; Inicia la creacion de las tablas necesarias.
	//	--recreateTables yes ; Recrea las tablas necesarias.
	//	--dropTables yes ; Borra las tablas necesarias.
	//	--port 8181	; Cambia el puerto por defecto del servidor.
	utils.FlagControl()

	//VerificationBasic Hace las verificaciones basicas para el funcionamiento del sistema.
	utils.VerificationBasic()

	//router Contiene una instancia ServerMux con todas las rutas ya definidas.
	router := routes.InitRoutes()

	//ServerSetup Ejecuta el servidor con el puerto y el router.
	utils.ServerSetup(utils.Port, router)
}
