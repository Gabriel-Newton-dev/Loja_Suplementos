package main

import (
	"net/http"
	"text/template"
)

type Produtos struct {
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	Produto := []Produtos{
		{"Creatina Universal", "A melhor creatina do mercado", 119.89, 10},
		{"HGH", "Otimizador de hormonio GH", 350.99, 5},
		{"Whey Protein", "O whey protein tem alta concentração de cálcio", 59.99, 10},
		{"Albumina", "Clara de ovo", 49.99, 50},
		{"Ômega 3", "Multivitaminíco", 32.99, 30},
		{"Natubolic", "Suplemento", 91.99, 4},
	}
	temp.ExecuteTemplate(w, "Index", Produto)
}
