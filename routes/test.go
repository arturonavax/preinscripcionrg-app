package routes

import (
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/controllers"
)

func SetTestRouter(mux *http.ServeMux) {
	prefix := "/test"

	mux.HandleFunc(prefix, controllers.Test)
}
