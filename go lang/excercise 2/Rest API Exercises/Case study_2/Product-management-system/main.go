package main

import (
	"log"
	"net/http"
	"product-management-system/controllers"
	"product-management-system/models"
)

func main() {
	models.InitDB()

	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/products/add", controllers.AddProduct)
	http.HandleFunc("/products/edit", controllers.EditProduct)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
