package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/databases"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/security"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/utils"
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

	log.Println("--> Un usuario esta intentando obtener un JSON Web Token. <--")

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
		log.Printf("$ UserID:%d { %s Solicito un JWT } $\n", userFound.ID, userFound.Username)

		//Se elimina la contraseña de la estructura por seguridad.
		userFound.Password = ""

		//Se genera un JWT con la estructura userFound
		token := security.GenerateJWT(userFound)

		//Se convierte el JWT resultante en un JSON valido para enviar al Usuario.
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir el token a JSON %s\n", err)
		}

		//Se envia el Token al ResponseWriter.
		w.WriteHeader(http.StatusAccepted)
		w.Write(j)

		log.Printf("$ UserID:%d { JWT Enviado a %s } $\n", userFound.ID, userFound.Username)
	} else {
		m := models.Message{
			Message: "Usuario o Clave no valido",
			Code:    http.StatusUnauthorized,
		}
		utils.DisplayMessage(w, m)
	}
}
