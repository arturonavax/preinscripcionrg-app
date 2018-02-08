package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println("No se encontro el archivo index.html")
	}

	t.Execute(w, nil)
}
