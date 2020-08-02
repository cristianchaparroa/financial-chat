package main

import (
	"chat/app/delivery/rest"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := rest.NewServer()
	s.Run()
}
