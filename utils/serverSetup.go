package utils

import (
	"fmt"
	"log"
	"net/http"
)

//ServerSetup Ejecuta el servidor con el puerto y el router.
func ServerSetup(port int, router *http.ServeMux) {
	log.Println("==========================")
	log.Println("| Ejecutando Servidor... |")
	log.Println("==========================")

	log.Println("================================================")
	log.Println(fmt.Sprintf("| Servidor escuchando := http://localhost:%d |", port))
	log.Println("================================================")
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", port),
			router,
		))

}
