package main

import (
	"net/http"
	configuration "visiku-restapi/Configuration"
	productcontroller "visiku-restapi/Controller/products"
)

func main() {
	// Call the DB
	configuration.DBConn()

	// REST API - Get & Post Product
	http.HandleFunc("/product", productcontroller.Product)

	// Turn on the service
	http.ListenAndServe(":8080", nil)
}
