package routes

import (
	"github.com/Gatusko/trafilea-http-numbers/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

type routes struct {
	numberController *controller.NumberController
}

func NewRoutes(numberController *controller.NumberController) *routes {
	return &routes{
		numberController: numberController,
	}
}

func (r *routes) MountAllRoutes() http.Handler {
	route := chi.NewRouter()
	// handling cors. For future, we should handle cors for each endpoint and what is neccesary
	route.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))
	route.Mount("/v1/numbers", r.numberController.NumberRoutes())
	return route
}
