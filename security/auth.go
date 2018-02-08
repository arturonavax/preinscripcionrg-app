package security

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/arthurnavah/PreInscripcionRG/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	//PublicKey contiene la llave Publica.
	PublicKey *rsa.PublicKey
)

//init Lee las llaves RSA.
func init() {
	var err error
	privateByte, _ := ioutil.ReadFile("./security/keys/private.rsa")

	publicBytes, _ := ioutil.ReadFile("./security/keys/public.rsa")

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		log.Println("No se pudo hacer el parse a privateKey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Println("No se pudo hacer el parse a PublicKey")
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
