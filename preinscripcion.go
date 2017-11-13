package main

import (
	"github.com/arthurnavah/PreInscripcionRG-API/routes"
	"github.com/arthurnavah/PreInscripcionRG-API/utils"
)

func main() {
	utils.FlagControl()

	router := routes.InitRoutes()

	utils.ServerSetup(utils.Port, router)
}
