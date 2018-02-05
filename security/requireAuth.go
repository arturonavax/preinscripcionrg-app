package security

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/arthurnavah/PreInscripcionRG-API/models"
	"github.com/arthurnavah/PreInscripcionRG-API/utils"
)

//RequireAuth Middleware para autenticar por JWT.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var m models.Message

		token, err := request.ParseFromRequestWithClaims(
			r,
			request.OAuth2Extractor,
			&models.Claim{},
			func(t *jwt.Token) (interface{}, error) {
				return PublicKey, nil
			},
		)

		if err != nil {
			m.Code = http.StatusUnauthorized

			switch err.(type) {
			case *jwt.ValidationError:
				vError := err.(*jwt.ValidationError)
				switch vError.Errors {
				case jwt.ValidationErrorExpired:
					m.Message = "Su token ha expirado"
					utils.DisplayMessage(w, m)
					return
				case jwt.ValidationErrorSignatureInvalid:
					m.Message = "La firma del token no coincide"
					utils.DisplayMessage(w, m)
					return
				default:
					m.Message = "Su token no es valido"
					utils.DisplayMessage(w, m)
					return
				}
			}
		}

		if token.Valid {
			ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			m.Code = http.StatusUnauthorized
			m.Message = "Su token no es valido"
			utils.DisplayMessage(w, m)

		}
	})
}
