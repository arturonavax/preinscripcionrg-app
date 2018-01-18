package routes

import (
	"net/http"

	"github.com/graphql-go/handler"
	"gitlab.com/arthurnavah/Production-Arthur/InventSoft2/schemas"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/security"
)

//SetAPIGraphql Declara la ruta del API Graphql.
func SetAPIGraphql(mux *http.ServeMux) {
	prefix := "/graphql"

	//Se crea el Handler GraphQL.
	hgql := handler.New(&handler.Config{
		Schema: &schemas.SchemaInit,
		Pretty: true,
	})

	//Se declara la ruta del API, El Middleware de Autenticacion con el Handler de GraphQL
	mux.Handle(prefix, security.RequireAuth(hgql))
}
