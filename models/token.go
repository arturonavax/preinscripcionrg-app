package models

//Token Estructura para devolver un Token al usuario Logeado.
type Token struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Token   string `json:"token,omitempty"`
}
