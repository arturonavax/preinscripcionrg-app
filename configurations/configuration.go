package configurations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//configurationDB Estructura que contiene la configuracion de la Base de datos
type configurationDB struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

//getConfigurationDB Obtiene la configuracion de un archivo JSON y pobla la estructura configurationDB
func getConfigurationDB() configurationDB {
	var c configurationDB
	file, err := os.Open("./configDB.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)

	return c
}

//GetConnectionDB Devuelve una conexion con la base de datos
func GetConnectionDB() *gorm.DB {
	c := getConfigurationDB()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Server, c.Port, c.Database, c.SSLMode)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
