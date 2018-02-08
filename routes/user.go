package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/controllers"
)

//SetRegisterRouter Declara la ruta para Registrar usuarios.
func SetUserRouter(mux *http.ServeMux) {
	prefix := "/api/users"
	prefixLogin := "/api/login"

	//Se declara el controlador de registracion para la ruta /api/users.
	mux.HandleFunc(prefix, controllers.Users)

	//Se declara el controlador de logeo para la ruta /api/login.
	mux.HandleFunc(prefixLogin, controllers.Login)
}
