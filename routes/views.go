package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/controllers"
)

//SetViewsRoutes Declara las rutas de las vistas con sus controladores.
func SetViewsRoutes(mux *http.ServeMux) {
	prefixRegister := "/register"
	mux.HandleFunc(prefixRegister, controllers.RegisterView)

	prefixLogin := "/login"
	mux.HandleFunc(prefixLogin, controllers.LoginView)

	prefixApp := "/app"
	mux.HandleFunc(prefixApp, controllers.AppView)
}
