package models

import "store-shop/db"

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

func InsertProduct(
	name string,
	description string,
	price float64,
	quantity int,
) {
	db := db.ConnectDatabase()

	insertProduct, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)
	defer db.Close()
}
