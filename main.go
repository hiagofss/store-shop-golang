package main

import (
	"net/http"
	"store-shop/routes"
)

func main() {
	routes.LoadRoutes()

	http.ListenAndServe("localhost:8000", nil)
}
