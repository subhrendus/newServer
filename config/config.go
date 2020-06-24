package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Logging LoggingConfig `yaml:"logging"`
	Service ServiceConfig `yaml:"service"`
}

type LoggingConfig struct {
	AppName    string `yaml:"app_name"`
	AppVersion string `yaml:"app_version"`
	Level      string `yaml:"level"`
}

type ServiceConfig struct {
	Listen int `yaml:"listen"`
}

// LoadConfiguration - read configurations from a yaml file and loads into 'Config' struct.
//					   returns error if the file is missing or contains bad schema.
func LoadConfiguration(filePath string) (*AppConfig, error) {
	cnf := &AppConfig{}

	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file %s", err.Error())
	}

	ymlErr := yaml.Unmarshal(rawContent, &cnf)
	if ymlErr != nil {
		return nil, fmt.Errorf("unable to unpack file %s", ymlErr.Error())
	}

	return cnf, nil
}