package controllers

import (
	"html/template"
	"net/http"
	"store-shop/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.InsertProduct(name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/products", http.StatusPermanentRedirect)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/products", http.StatusPermanentRedirect)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.GetProduct(productId)

	temp.ExecuteTemplate(w, "EditProduct", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err.Error())
		}
		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(id, name, description, priceConverted, quantityConverted)
	}

	http.Redirect(w, r, "/products", http.StatusPermanentRedirect)
}
