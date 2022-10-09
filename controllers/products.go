package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"store-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")
		priceFormatted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println(err)
		}
		amountFormatted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println(err)
		}
		models.InsertProduct(name, description, priceFormatted, amountFormatted)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.GetProduct(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")
		idFormatted, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
		}
		priceFormatted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println(err)
		}
		amountFormatted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println(err)
		}
		models.UpdateProduct(idFormatted, name, description, priceFormatted, amountFormatted)
	}
	http.Redirect(w, r, "/", 301)
}
