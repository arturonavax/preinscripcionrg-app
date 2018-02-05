package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG-API/controllers"
)

func SetTestRouter(mux *http.ServeMux) {
	prefix := "/test"

	mux.HandleFunc(prefix, controllers.Test)
}
