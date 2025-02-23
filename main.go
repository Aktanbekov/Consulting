package main

import (
	"aktai/handlers"
	"aktai/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":8080")
}