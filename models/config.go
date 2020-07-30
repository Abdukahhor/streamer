package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Config - configuration of the service
type Config struct {
	URLs             []string `yaml:"URLs"`
	MinTimeout       int      `yaml:"MinTimeout"`
	MaxTimeout       int      `yaml:"MaxTimeout"`
	NumberOfRequests int      `yaml:"NumberOfRequests"`
	Ln               int
}

//Get -
func (c *Config) Get(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, c)
	if err != nil {
		return err
	}
	return nil
}
