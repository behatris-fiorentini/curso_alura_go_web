package main

import (
	"github.com/behatris-fiorentini/curso_alura_go_web/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
