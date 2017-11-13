package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim Contenedor del JWT
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
