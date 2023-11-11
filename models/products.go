package models

import "github.com/hiagofss/store-shop/db"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {

	db := db.ConnectDatabase()
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
	return products
}
