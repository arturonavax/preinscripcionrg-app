package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RegisterView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./public/static/register.html")
		if err != nil {
			fmt.Println("No se encontro el arrchivo register.html")
		}

		t.Execute(w, nil)
	}
}
