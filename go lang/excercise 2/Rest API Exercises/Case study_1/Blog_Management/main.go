package main

import (
	"log"
	"net/http"

	"project_root/config"
	"project_root/handlers"
	"project_root/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	router := mux.NewRouter()

	router.Use(middleware.Logging)
	router.Use(middleware.Validation)
	router.Use(middleware.Authentication)

	router.HandleFunc("/blog", handlers.CreateBlog).Methods("POST")
	router.HandleFunc("/blog/{id}", handlers.GetBlog).Methods("GET")
	router.HandleFunc("/blogs", handlers.GetBlogs).Methods("GET")
	router.HandleFunc("/blog/{id}", handlers.UpdateBlog).Methods("PUT")
	router.HandleFunc("/blog/{id}", handlers.DeleteBlog).Methods("DELETE")

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
