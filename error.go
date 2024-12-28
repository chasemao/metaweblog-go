package main

import "errors"

var (
	errInvalidTitle       = errors.New("invalid title")
	errInvalidDescription = errors.New("invalid description")
	errInvalidUserName    = errors.New("invalid user name")
	errInvalidPassword    = errors.New("invalid password")
)
