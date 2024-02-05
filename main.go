package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/product-manager/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
