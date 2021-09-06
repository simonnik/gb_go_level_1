package main

import (
	"./config"
	"log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Println("Error reading config.")
		return
	}
	log.Println("Config properties:")
	log.Println("JaegerUrl: ", conf.JaegerUrl)
	log.Println("Port: ", conf.Port)
	log.Println("DbUrl: ", conf.DbUrl)
	log.Println("KafkaBroker: ", conf.KafkaBroker)
}
