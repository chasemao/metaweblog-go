package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var c config

type config struct {
	BlogURL                string `yaml:"blogURL"`
	BlogTitle              string `yaml:"blogTitle"`
	BlogDir                string `yaml:"blogDir"`
	MediaDir               string `yaml:"mediaDir"`
	MediaRelDirForBlogHtml string `yaml:"mediaRelDirForBlogHtml"`
	UserName               string `yaml:"userName"`
	Password               string `yaml:"password"`
	HugoCompileDir         string `yaml:"hugoCompileDir"` //TODO
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
