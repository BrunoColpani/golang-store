package controllers

import (
	"log"
	"net/http"
	"products/models"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFormattedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}
		quantityFormattedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CreateProduct(name, description, priceFormattedToFloat, quantityFormattedToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idFormattedToInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do id para int", err)
		}

		quantityFormattedToInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade para int", err)
		}

		priceFormattedToFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço para float", err)
		}

		models.UpdateProduct(idFormattedToInt, name, description, priceFormattedToFloat, quantityFormattedToInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
