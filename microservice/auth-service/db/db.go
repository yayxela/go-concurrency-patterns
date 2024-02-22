package db

import "errors"

var (
	NotExists     = errors.New("not exists")
	AlreadyExists = errors.New("already exist")
)
