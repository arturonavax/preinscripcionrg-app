package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG-API/databases"
	"github.com/arthurnavah/PreInscripcionRG-API/models"
	"github.com/arthurnavah/PreInscripcionRG-API/utils"
)

//Login Controlador de Logeo.
func Login(w http.ResponseWriter, r *http.Request) {
	userFound := models.User{}

	//Se decodifica el cuerpo del Request en userFound.
	err := json.NewDecoder(r.Body).Decode(&userFound)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	//Se instancia una conexion a la base de datos.
	db := databases.GetConnectionDB()
	defer db.Close()

	//Se encripta con sha256 la contraseña enviada en el cuerpo del Request.
	c := sha256.Sum256([]byte(userFound.Password))
	pwd := fmt.Sprintf("%x", c)

	//Se pobla userFound con el usuario que cumpla el Query.
	db.Where("username = ? AND password = ?", userFound.Username, pwd).First(&userFound)

	//userFound.ID es mayor a 0 cuando tuvo algun resultado el Query.
	if userFound.ID > 0 {
		//Se elimina la contraseña de la estructura por seguridad.
		userFound.Password = ""

		//Se genera un JWT con la estructura userFound
		token := utils.GenerateJWT(userFound)

		//Se convierte el JWT resultante en un JSON valido para enviar al Usuario.
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a JSON %s\n", err)
		}

		//Se envia el Token al ResponseWriter.
		w.WriteHeader(http.StatusAccepted)
		w.Write(j)

	} else {
		m := models.Message{
			Message: "Usuario o Clave no valido",
			Code:    http.StatusUnauthorized,
		}
		utils.DisplayMessage(w, m)
	}
}
