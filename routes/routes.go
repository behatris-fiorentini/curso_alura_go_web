package routes

import (
	"github.com/behatris-fiorentini/curso_alura_go_web/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/store", controllers.Store)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
