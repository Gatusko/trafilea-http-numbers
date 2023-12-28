package model

type Error struct {
	Message string `json:"error"`
}

func NewError(msg string) *Error {
	return &Error{Message: msg}
}
