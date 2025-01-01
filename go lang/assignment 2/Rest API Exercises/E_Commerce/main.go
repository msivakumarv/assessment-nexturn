package main

import (
	"ecommerce-inventory-service/config"
	"ecommerce-inventory-service/handlers"
	"ecommerce-inventory-service/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()

	router := mux.NewRouter()

	router.Use(handlers.LoggingMiddleware)
	router.Use(handlers.AuthMiddleware)

	routes.RegisterProductRoutes(router)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
