package controllers

import (
	"html/template"
	"net/http"
	"store-shop/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}
