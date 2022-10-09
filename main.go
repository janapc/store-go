package main

import (
	"net/http"

	"store-go/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
