package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Gabriel-Newton-dev/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		valor := r.FormValue("valor")
		quantidade := r.FormValue("quantidade")

		valorConvertido, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			log.Println("Erro na conversão do valor:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, valorConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func DeletarNoBanco(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

// Update Busca as informacoes do dormulário.
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		valor := r.FormValue("valor")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID", err)
		}

		valorConvertidoParaFloat, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			log.Println("Erro na conversão do valor", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão de quantidade", err)
		}

		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, valorConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
