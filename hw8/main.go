package main

import (
	"github.com/simonnik/gb_go_level_1/hw8/config"
	"log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Println("Error reading config.", err)
		return
	}
	log.Println("Config properties:")
	log.Println("JaegerUrl: ", conf.JaegerUrl)
	log.Println("Port: ", conf.Port)
	log.Println("DbUrl: ", conf.DbUrl)
	log.Println("KafkaBroker: ", conf.KafkaBroker)
}
