package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/controllers"
)

func SetViewsRoutes(mux *http.ServeMux) {
	prefixRegister := "/register"
	mux.HandleFunc(prefixRegister, controllers.RegisterView)

	prefixLogin := "/login"
	mux.HandleFunc(prefixLogin, controllers.LoginView)
}
