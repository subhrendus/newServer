package main

import (
	"flag"
)

func main() {
	var configPath string

	//read the value for config into configPath var
	//if no value is given, 'config/config.yml' is the default
	flag.StringVar(&configPath, "config", "config/config.yml", "The location for the config YAML file.")

	api := App{}
	api.Initialize(configPath)
	api.Run()	
}