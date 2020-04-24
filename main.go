package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	server_host    string `yaml:"server_host"`
	query_port     int    `yaml:"query_port"`
	sid            int    `yaml:"sid"`
	query_login    string `yaml:"query_login"`
	query_password string `yaml:"query_password"`
}

func ConfigInit(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new yaml decode
	d := yaml.NewDecoder(file)

	// Start yaml decodint from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	fmt.Println("vim-go")
}
