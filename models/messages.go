package models

//Message Mensaje para el usuario.
type Message struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

//CreateMessage Mensaje de creacion para el usuario.
type CreateMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
