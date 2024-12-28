package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var c config

type config struct {
	BlogURL     string `yaml:"blogURL"`
	BlogTitle   string `yaml:"blogTitle"`
	HugoRootDir string `yaml:"hugoRootDir"`
	UserName    string `yaml:"userName"`
	Password    string `yaml:"password"`
}

func initConfig() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(c)
	return nil
}

func getConfig() config {
	return c
}