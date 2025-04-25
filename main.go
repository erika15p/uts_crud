package main

import (
	"UTS/database"
	"UTS/handlers"
	"net/http"
)

func main() {
	database.ConnectDB()

	http.HandleFunc("/categories", handlers.GetCategories)
	http.HandleFunc("/categories/create", handlers.CreateCategory)

	http.HandleFunc("/products", handlers.GetProducts)
	http.HandleFunc("/products/create", handlers.CreateProduct)

	http.HandleFunc("/products/update", handlers.UpdateProduct)
	http.HandleFunc("/products/delete", handlers.DeleteProduct)

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/main.js", http.StripPrefix("/", http.FileServer(http.Dir("./static/js"))))

	http.ListenAndServe(":8080", nil)
}
