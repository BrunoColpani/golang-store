package main

import (
	"net/http"
	"products/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
