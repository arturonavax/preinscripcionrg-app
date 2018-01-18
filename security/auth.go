package security

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/models"
)

var (
	privateKey *rsa.PrivateKey
	//PublicKey contiene la llave Publica.
	PublicKey *rsa.PublicKey
)

//init Lee las llaves RSA.
func init() {
	privateByte, err := ioutil.ReadFile("./security/keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}

	publicBytes, err := ioutil.ReadFile("./security/keys/public.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a PublicKey")
	}
}

//GenerateJWT Genera un JSON Web Token.
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer: "PimoSoft",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}

	return result
}
