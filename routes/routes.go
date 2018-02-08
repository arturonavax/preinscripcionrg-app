package routes

import "net/http"

//InitRoutes Devuelve un ServeMux con todas las rutas ya definidas.
func InitRoutes() *http.ServeMux {
	//Se crea una instancia ServeMux.
	mux := http.NewServeMux()

	//Se Establece la ruta /api/users y /api/login
	SetUserRouter(mux)

	//Se Establece la ruta /graphql
	SetAPIGraphql(mux)

	//Se Expone la carpeta publica.
	SetPublicRoutes(mux)

	//Se Establecen las rutas de las vistas.
	SetViewsRoutes(mux)

	//Se devuelve la instacia ServeMux con las rutas declaradas
	return mux
}
