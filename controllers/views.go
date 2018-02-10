package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

//RegisterView Controlador para la ruta /register.
func RegisterView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./public/static/register.html")
		if err != nil {
			fmt.Println("No se encontro el arrchivo static/register.html")
		}

		t.Execute(w, nil)
	}
}

//LoginView Controlador para la ruta /login.
func LoginView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println("No se encontro el archivo public/index.html")
	}

	t.Execute(w, nil)
}

//AppView Controlador para la ruta /app.
func AppView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public/static/app/profile.html")
	if err != nil {
		fmt.Println("No se encontro el archivo app/profile.html")
	}

	t.Execute(w, nil)
}
