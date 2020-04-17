package main

import (
	"net/http"
	"text/template"

	"github.com/luis/crud/routes"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
