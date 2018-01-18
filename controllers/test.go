package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
)

func Test(w http.ResponseWriter, r *http.Request) {
	prueba := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&prueba)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	prueba.Message += "RESPUESTA"

	j, _ := json.Marshal(prueba)
	w.WriteHeader(prueba.Code)
	w.Write(j)
}
