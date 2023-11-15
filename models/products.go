package models

import (
	"store-shop/db"
)

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

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func GetProduct(id string) Product {
	db := db.ConnectDatabase()

	productList, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productList.Next() {
		var id, name, description string
		var quantity int
		var price float64

		err := productList.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()
	return product
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

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func UpdateProduct(
	id string,
	name string,
	description string,
	price float64,
	quantity int,
) {
	db := db.ConnectDatabase()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
