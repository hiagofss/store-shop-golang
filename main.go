package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/app?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/products", renderProducts)
	http.ListenAndServe("localhost:8000", nil)
}

func renderProducts(w http.ResponseWriter, r *http.Request) {
	db := connectDatabase()
	productsList, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}

	product := Product{}
	products := []Product{}

	for productsList.Next() {
		var id, name, description string
		var quantity int
		var price float64

		err := productsList.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err)
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()
	temp.ExecuteTemplate(w, "Index", products)
}
