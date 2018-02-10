package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/security"
	"github.com/arthurnavah/PreInscripcionRG/utils"
)

//Users Controlador para las rutas /api/users.
func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		UserCreate(w, r)
	}
}

//UserCreate Controlador de Registro de Usuario.
func UserCreate(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(user)

	//user.Password y user.ConfirmPassword deben coincidir para continuar.
	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusNotAcceptable
		utils.DisplayMessage(w, m)
		return

	} else if user.Password == user.ConfirmPassword {

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
			log.Printf("-+> Se a registrado una Persona : '%s %s' (PeopleID:%d) <+-\n",
				people.FirstName,
				people.LastName,
				people.ID,
			)

			user.PeopleID = people.ID

		}

		//Se crea el usuario en la base de datos
		err = tx.Create(&user).Error
		if err != nil {
			m.Message = fmt.Sprintf("Error al crear el usuario: %s", err)
			m.Code = http.StatusBadRequest
			utils.DisplayMessage(w, m)
			tx.Rollback()
			return

		} else {
			//Transaccion exitosa
			tx.Commit()

			log.Printf("-+> Se a registrado un Usuario : '%s' (UserID:%d) (PeopleID:%d) <+-\n",
				user.Username,
				user.ID,
				people.ID,
			)

			//Se envia el mensaje de exito al usuario.
			m2.Message = "Usuario creado con Exitoo"
			m2.Code = http.StatusCreated
			m2.ID = user.ID
			utils.DisplayCreateMessage(w, m2)

		}
	}

}

// --------------------------------------------------------------------------------

//Login Controlador de Logeo en la ruta /api/login.
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
		log.Printf("$ UserID:%d { %s Solicito un JWT } $\n",
			userFound.ID,
			userFound.Username,
		)

		//Se elimina la contraseña de la estructura por seguridad.
		userFound.Password = ""

		//Se genera un JWT con la estructura userFound
		token := security.GenerateJWT(userFound)

		//Se convierte el JWT resultante en un JSON valido para enviar al Usuario.
		j, err := json.Marshal(models.Token{
			Token:   token,
			Message: "Logueo Exitoso!",
			Code:    http.StatusContinue,
		})
		if err != nil {
			log.Fatalf("Error al convertir el token a JSON %s\n", err)
		}

		//Se envia el Token al ResponseWriter.
		w.WriteHeader(http.StatusAccepted)
		w.Write(j)

		log.Printf("$ UserID:%d { JWT Enviado a %s } $\n",
			userFound.ID,
			userFound.Username,
		)

	} else {
		j, _ := json.Marshal(models.Token{
			Token:   "",
			Message: "Usuario o Contraseña invalido",
			Code:    http.StatusNonAuthoritativeInfo,
		})

		w.Write(j)

	}

}
