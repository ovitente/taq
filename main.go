package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	ServerHost    string `yaml:"server_host"`
	QueryPort     int64  `yaml:"query_port"`
	Sid           int64  `yaml:"sid"`
	QueryLogin    string `yaml:"query_login"`
	QueryPassword string `yaml:"query_password"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	fmt.Println(c)

}
