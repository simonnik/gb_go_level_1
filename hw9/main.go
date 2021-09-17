package main

import (
	"flag"
	"github.com/simonnik/gb_go_level_1/hw9/config"
	"log"
)

func main() {
	var configFile = flag.String("file", "", "Path to yaml or json config file")

	flag.Parse()
	conf, err := config.Load(configFile)
	if err != nil {
		log.Println("Error reading config.", err)
		return
	}
	log.Println("Config properties:")
	log.Println("Host: ", conf.Host)
	log.Println("Port: ", conf.Port)
	log.Println("Timeout: ", conf.Timeout)
}
