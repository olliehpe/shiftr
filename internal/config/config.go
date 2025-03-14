package config

import (
	"github.com/olliehpe/shiftr/internal"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func LoadConfig() *internal.Config {
	var config *internal.Config
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Printf("error reading config.yml: %v", err)
	}
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("error parsing config.yml: %v", err)
	}
	return config
}
