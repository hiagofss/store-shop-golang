package routes

import (
	"net/http"
	"store-shop/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/products", http.StatusMovedPermanently)
	})

	http.HandleFunc("/products", controllers.GetProducts)
	http.HandleFunc("/products/new", controllers.NewProduct)
	http.HandleFunc("/products/insert", controllers.InsertProduct)
}
