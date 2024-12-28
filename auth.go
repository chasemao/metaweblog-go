package main

import "fmt"

func auth(userName string, password string) error {
	if getConfig().UserName != "" && userName != getConfig().UserName {
		fmt.Println(errInvalidUserName)
		return errInvalidUserName
	}
	if getConfig().Password != "" && password != getConfig().Password {
		fmt.Println(errInvalidPassword)
		return errInvalidPassword
	}
	return nil
}
