package api

import "errors"

type ErrResponse struct {
	Error string `json:"error"`
}

var (
	WrongEmailOrPassword = errors.New("wrong email or password")
)
