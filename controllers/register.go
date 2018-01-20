package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/databases"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/utils"
)

//Register Controlador de Registro de Usuario.
func Register(w http.ResponseWriter, r *http.Request) {
	people := models.People{}
	user := models.User{}
	m := models.Message{}
	m2 := models.CreateMessage{}

	//Se decodifica el cuerpo del Request en user.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a Registrar : %s", err)
		m.Code = http.StatusBadRequest
		utils.DisplayMessage(w, m)
		return
	}

	//user.Password y user.ConfirmPassword deben coincidir para continuar.
	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		utils.DisplayMessage(w, m)
		return
	}

	//Se instancia una conexion a la base de datos.
	db := databases.GetConnectionDB()
	defer db.Close()

	//Se encripta con sha256 la contraseña enviada en el cuerpo del Request.
	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)

	//Se establece la contraseña encriptada el Password de user.
	user.Password = pwd

	//Poblando la persona
	people.FirstName = user.FirstName
	people.LastName = user.LastName
	people.Email = user.Email
	people.Address = user.Address
	people.PhoneNumber = user.PhoneNumber

	// Inicio de la Transaccion
	tx := db.Begin()

	//Se crea la persona en la base de datos en la base de datos.
	err = tx.Create(&people).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear la persona: %s", err)
		m.Code = http.StatusBadRequest
		utils.DisplayMessage(w, m)
		tx.Rollback()
		return
	} else {
		log.Printf("-+> Se a registrado una Persona : '%s %s' (PeopleID:%d) <+-\n", people.FirstName, people.LastName, people.ID)
	}

	//Se crea el usuario en la base de datos
	user.PeopleID = people.ID
	err = tx.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("Error al crear el usuario: %s", err)
		m.Code = http.StatusBadRequest
		utils.DisplayMessage(w, m)
		tx.Rollback()
		return
	} else {
		log.Printf("-+> Se a registrado un Usuario : '%s' (UserID:%d) (PeopleID:%d) <+-\n", user.Username, user.ID, people.ID)
	}

	tx.Commit()
	//Se envia el mensaje de exito al usuario.
	m2.Message = "Usuario creado con Exitoo"
	m2.Code = http.StatusCreated
	m2.ID = user.ID
	utils.DisplayCreateMessage(w, m2)
}
