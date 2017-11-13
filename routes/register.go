package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG-API/controllers"
)

//SetRegisterRouter Declara la ruta para Registrar usuarios.
func SetRegisterRouter(mux *http.ServeMux) {
	prefix := "/register"

	//Se declara el controlador de registracion para la ruta /register.
	mux.HandleFunc(prefix, controllers.Register)
}
