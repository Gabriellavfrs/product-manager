package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/product-manager/models"
)

// variável para armazenar todos os templates
// Must encapsula todos os templates de devolve 2 retornos
var temp = template.Must(template.ParseGlob("templates/*html"))

// w é quem consegue passar a resposta
func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetAllProducts()

	// (quem passa a resposta, quem vamos exibir, informações que queremos passar para a página)
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// os valores virão como string
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		amount := r.FormValue("quantidade")

		formatedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		formatedAmount, err := strconv.Atoi(amount)

		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.InsertProduct(name, description, formatedPrice, formatedAmount)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Edit", nil)

}
