package routes

import (
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/controllers"
)

//SetLoginRouter Declara la ruta de Logeo del usuario.
func SetLoginRouter(mux *http.ServeMux) {
	prefix := "/login"

	//Se declara el controlador de Logeo para la ruta /login.
	mux.HandleFunc(prefix, controllers.Login)
}
