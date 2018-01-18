package routes

import "net/http"

//InitRoutes Devuelve un ServeMux con todas las rutas ya definidas.
func InitRoutes() *http.ServeMux {
	//Se crea una instancia ServeMux.
	mux := http.NewServeMux()

	//Se Establece la ruta /register
	SetRegisterRouter(mux)

	//Se Establece la ruta /login
	SetLoginRouter(mux)

	//Se Establece la ruta /graphql

	//Prueba
	SetTestRouter(mux)

	//Se devuelve la instacia ServeMux con las rutas declaradas
	return mux
}
