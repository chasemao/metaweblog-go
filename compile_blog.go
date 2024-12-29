package main

import (
	"log"
	"os/exec"
)

//lint:ignore U1000 todo part
func (a *apiImpl) compileBlog() {
	if getConfig().HugoCompileDir == "" {
		return
	}
	cmd := exec.Command("hugo", "--gc", "--minify", "-s", getConfig().HugoCompileDir)
	output, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(output))
}
