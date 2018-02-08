package routes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/schemas"
	"github.com/arthurnavah/PreInscripcionRG/security"
	"github.com/graphql-go/handler"
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
