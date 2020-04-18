package main

import (
	"net/http"
	"github.com/luis/crud/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
