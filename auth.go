package main

import "fmt"

func auth(userName string, password string) error {
	if getConfig().UserName != "" && userName != getConfig().UserName {
		fmt.Println(invalidUserNameError)
		return invalidUserNameError
	}
	if getConfig().Password != "" && password != getConfig().Password {
		fmt.Println(invalidPasswordError)
		return invalidPasswordError
	}
	return nil
}
