package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
)

//DisplayMessage Envia un mensaje al Usuario.
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatal("Error al convertir el mensaje:", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}

//DisplayCreateMessage Envia un mensaje de creacion al usuario al Usuario.
func DisplayCreateMessage(w http.ResponseWriter, m models.CreateMessage) {
	j, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatal("Error al convertir el mensaje:", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
