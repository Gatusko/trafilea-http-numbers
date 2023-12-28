package main

import (
	"fmt"
	"github.com/Gatusko/trafilea-http-numbers/controller"
	"github.com/Gatusko/trafilea-http-numbers/domain/repositories"
	"github.com/Gatusko/trafilea-http-numbers/routes"
	"github.com/Gatusko/trafilea-http-numbers/services"
	"log"
	"net/http"
	"os"
)

// Starting point of the server
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Print("Using default port 8080")
		port = "8080"
	}

	// Starting the Dependecy  Injection
	numberRepository := repositories.NewNumberMemoryRepository()
	numberService := services.NewNumberSerivce(numberRepository)
	numberController := controller.NewNumberController(numberService)
	routes := routes.NewRoutes(numberController)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), routes.MountAllRoutes())
	if err != nil {
		log.Fatalf("Error starting http server %v", err)
	}
	fmt.Println(port)
}
