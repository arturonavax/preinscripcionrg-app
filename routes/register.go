package routes

import (
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/controllers"
)

//SetRegisterRouter Declara la ruta para Registrar usuarios.
func SetRegisterRouter(mux *http.ServeMux) {
	prefix := "/register"

	//Se declara el controlador de registracion para la ruta /register.
	mux.HandleFunc(prefix, controllers.Register)
}
