package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	server_host    string `yaml:"server_host"`
	query_port     int64  `yaml:"query_port"`
	sid            int64  `yaml:"sid"`
	query_login    string `yaml:"query_login"`
	query_password string `yaml:"query_password"`
}

func (c *config) getConf() *config {
	configPath := "./config.yml"
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("Config file not found #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c

}

func main() {
	var c config
	c.getConf()

	fmt.Println(c)

}
