package main

import (
	"log"

	"github.com/RafaelCavalcante17/api-students/api"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
