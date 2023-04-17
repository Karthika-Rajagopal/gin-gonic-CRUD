package main

import (
	"e/handlers"
	"e/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}