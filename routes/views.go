package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/controllers"
)

//SetViewsRoutes Declara las rutas de las vistas con sus controladores.
func SetViewsRoutes(mux *http.ServeMux) {
	prefixUser := "/user"
	mux.HandleFunc(prefixUser, controllers.UserView)

	prefixApp := "/app"
	mux.HandleFunc(prefixApp, controllers.AppView)

	prefixPlantilla := "/plantilla"
	mux.HandleFunc(prefixPlantilla, controllers.PlantillaView)
}
