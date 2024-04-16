package tester

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)
type Configuration struct {
	SuccessConditionsConfiguration []SuccessConditionConfiguration `yaml:"success_conditions"`
}

type SuccessConditionConfiguration struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func ParseYMLToConfig(filepath string) Configuration {
	f, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		log.Fatal("Error reading configuration yaml. Aborting")
	}

	var c Configuration

	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Println(err)
		log.Fatal("Error reading configuration yaml. Aborting")
	}

	return c
}