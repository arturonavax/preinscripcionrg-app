package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

//UserView Controlador para la ruta /user.
func UserView(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Estan haciendo un Request a la pagina user.html")
	t, err := template.ParseFiles("./public/static/user.html")
	if err != nil {
		fmt.Println("No se encontro el archivo static/user.html")
	}

	t.Execute(w, nil)
}

//AppView Controlador para la ruta /app.
func AppView(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Estan haciendo un Request a la pagina main.html")
	t, err := template.ParseFiles("./public/static/app/main.html")
	if err != nil {
		fmt.Println("No se encontro el archivo app/main.html")
	}

	t.Execute(w, nil)
}
