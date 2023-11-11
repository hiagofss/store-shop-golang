package main

import (
	"net/http"
	"text/template"

	"github.com/hiagofss/store-shop/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/products", renderProducts)
	http.ListenAndServe("localhost:8000", nil)
}

func renderProducts(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}
