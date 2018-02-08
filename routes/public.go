package routes

import "net/http"

//SetPublicRoutes Expone la carpeta publica del servidor.
func SetPublicRoutes(mux *http.ServeMux) {
	dir := http.Dir("./public")
	fileServer := http.FileServer(dir)
	mux.Handle("/", fileServer)
}
