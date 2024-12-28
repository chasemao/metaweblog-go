package main

import "errors"

var (
	invalidTitleError       = errors.New("invalid title")
	invalidDescriptionError = errors.New("invalid description")
	invalidUserNameError    = errors.New("invalid user name")
	invalidPasswordError    = errors.New("invalid password")
	invalidMIMEType         = errors.New("invalid mime type")
)
