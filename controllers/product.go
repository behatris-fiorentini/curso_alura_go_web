package controllers

import (
	"github.com/behatris-fiorentini/curso_alura_go_web/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products, _ := models.GetProduct()
	temp.ExecuteTemplate(w, "Index", products)
}

func Store(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Store", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			panic(err.Error())
		}
		quantity, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
		if err != nil {
			panic(err.Error())
		}

		models.Create(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	models.Delete(ID)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	product := models.GetByID(ID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		panic(err.Error())
	}
	quantity, err := strconv.ParseInt(r.FormValue("quantity"), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	models.Update(id, name, description, price, quantity)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
