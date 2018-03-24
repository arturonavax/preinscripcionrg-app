package databases

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

//ConfigurationDB Estructura que contiene la configuracion de la Base de datos
type ConfigurationDB struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

//getConfigurationDB Obtiene la configuracion de un archivo JSON y pobla la estructura configurationDB
func getConfigurationDB() ConfigurationDB {
	var c ConfigurationDB
	file, err := os.Open("./databases/configDB.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)

	return c
}

//GetConnectionDB Devuelve una conexion con la base de datos
func GetConnectionDB() *gorm.DB {
	/*
		c := getConfigurationDB()

		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			c.User, c.Password, c.Server, c.Port, c.Database, c.SSLMode)

		db, err := gorm.Open("postgres", dsn)
		if err != nil {
			log.Fatal(err)
		}

		return db
	*/
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	return db

}
