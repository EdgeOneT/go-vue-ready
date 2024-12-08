package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

var Cfg Config

func LoadConfig(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Cfg); err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
